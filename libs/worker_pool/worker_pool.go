package workerpool

import (
	"context"
	"sync"
)

type WorkerFunc[In, Out any] func(ctx context.Context, arg In) (Out, error)

type Task[In, Out any] struct {
	fn  WorkerFunc[In, Out]
	arg In
}

type Result[Out any] struct {
	Value Out
	Err   error
}

type Pool[In, Out any] struct {
	maxWorkers int

	taskChan chan Task[In, Out]
	resChan  chan Result[Out]
	results  []Result[Out]

	workWg   sync.WaitGroup
	resultWg sync.WaitGroup

	closeWorking sync.Once
	closeRes     sync.Once
}

func NewPool[In, Out any](ctx context.Context, maxWorkers int) *Pool[In, Out] {
	if maxWorkers < 1 {
		maxWorkers = 1
	}

	p := &Pool[In, Out]{
		maxWorkers: maxWorkers,
		taskChan:   make(chan Task[In, Out]),
		resChan:    make(chan Result[Out]),
	}

	go p.run(ctx)

	return p
}

func (p *Pool[In, Out]) SubmitOne(fn WorkerFunc[In, Out], task In) {
	p.taskChan <- Task[In, Out]{fn, task}
}

func (p *Pool[In, Out]) SubmitMany(fn WorkerFunc[In, Out], tasks []In) {
	for _, task := range tasks {
		p.SubmitOne(fn, task)
	}
}

func (p *Pool[In, Out]) GetResult() []Result[Out] {
	return p.results
}

func (p *Pool[In, Out]) run(ctx context.Context) {
	for i := 0; i < p.maxWorkers; i++ {
		p.workWg.Add(1)

		go work(ctx, &p.workWg, p.taskChan, p.resChan)
	}

	p.resultWg.Add(1)

	go p.accumulateResult()
}

func (p *Pool[In, Out]) accumulateResult() {
	defer p.resultWg.Done()

	for res := range p.resChan {
		p.results = append(p.results, res)
	}
}

func work[In, Out any](
	ctx context.Context,
	wg *sync.WaitGroup,
	taskChan <-chan Task[In, Out],
	resChan chan<- Result[Out],
) {
	defer wg.Done()

	for {
		task, ok := <-taskChan
		if !ok {
			// If the channel is closed, exit the loop.
			break
		}

		select {
		case <-ctx.Done():
			break
		default:
			// Execute the task and send the result to resChan.
			value, err := task.fn(ctx, task.arg)
			resChan <- Result[Out]{value, err}
		}
	}
}

func (p *Pool[In, Out]) Wait() {
	p.closeWorking.Do(func() {
		close(p.taskChan)
	})

	p.workWg.Wait()

	p.closeRes.Do(func() {
		close(p.resChan)
	})

	p.resultWg.Wait()
}