package rate_limiter

import (
	"context"
	"sync"
	"time"
)

type SlidingWindow struct {
	limit    int
	interval time.Duration
	lastTime time.Time

	mu sync.Mutex

	prevCount int
	curCount  int
}

func NewSlidingWindow(limit int, interval time.Duration) *SlidingWindow {
	return &SlidingWindow{
		limit:    limit,
		interval: interval,
		lastTime: time.Now(),
	}
}

func (rl *SlidingWindow) Allow(_ context.Context) (bool, error) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastTime)

	if elapsed >= rl.interval {
		rl.prevCount = rl.curCount
		rl.curCount = 0
		rl.lastTime = now
	}

	slidingWindowCount := (float64(rl.prevCount) * (rl.interval.Seconds() - elapsed.Seconds())) / rl.interval.Seconds()

	curCount := rl.curCount + int(slidingWindowCount)

	if curCount >= rl.limit {
		return false, nil
	}

	rl.curCount++

	return true, nil
}

func (rl *SlidingWindow) Wait(ctx context.Context) error {
	for {
		allow, err := rl.Allow(ctx)
		if err != nil {
			return err
		}

		if allow {
			break
		}

		// Because we have a sliding window, we will wait for a half of the period and retry
		time.Sleep(rl.interval / 2)
	}

	return nil
}
