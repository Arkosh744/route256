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
	workers int

	taskCh   chan Task[In, Out]
	resultCh chan Result[Out]
	results  []Result[Out]

	taskWg   sync.WaitGroup
	resultWg sync.WaitGroup

	closeWorking sync.Once
	closeRes     sync.Once
}

func NewPool[In, Out any](ctx context.Context, maxWorkers int) *Pool[In, Out] {
	if maxWorkers < 1 {
		maxWorkers = 1
	}

	p := &Pool[In, Out]{
		workers:  maxWorkers,
		taskCh: make(chan Task[In, Out]),
		resultCh:  make(chan Result[Out]),
	}

	go p.run(ctx)

	return p
}

func (p *Pool[In, Out]) SendOne(fn WorkerFunc[In, Out], arg In) {
	p.taskCh <- Task[In, Out]{fn, arg}
}

func (p *Pool[In, Out]) SendMany(fn WorkerFunc[In, Out], args []In) {
	for _, task := range args {
		p.SendOne(fn, task)
	}
}

func (p *Pool[In, Out]) GetResult() []Result[Out] {
	return p.results
}

func (p *Pool[In, Out]) run(ctx context.Context) {
	for i := 0; i < p.workers; i++ {
		p.taskWg.Add(1)

		go work(ctx, &p.taskWg, p.taskCh, p.resultCh)
	}

	p.resultWg.Add(1)
	// reading results in separate goroutine until taskCh is closed.
	go p.getResults()
}

func (p *Pool[In, Out]) getResults() {
	defer p.resultWg.Done()

	for res := range p.resultCh {
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
			// Execute the task and send the result to resultCh.
			value, err := task.fn(ctx, task.arg)
			resChan <- Result[Out]{value, err}
		}
	}
}

func (p *Pool[In, Out]) Wait() {
	// close taskCh to stop all workers and wait for them to finish.
	p.closeWorking.Do(func() {
		close(p.taskCh)
	})

	p.taskWg.Wait()

	// close resultCh to stop getResults goroutine and wait for it.
	p.closeRes.Do(func() {
		close(p.resultCh)
	})

	p.resultWg.Wait()
}
