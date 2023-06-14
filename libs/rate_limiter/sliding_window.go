package rate_limiter

import (
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

func (l *SlidingWindow) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(l.lastTime)

	if elapsed >= l.interval {
		l.prevCount = l.curCount
		l.curCount = 0
		l.lastTime = now
	}

	slidingWindowCount := (float64(l.prevCount) * (l.interval.Seconds() - elapsed.Seconds())) / l.interval.Seconds()

	curCount := l.curCount + int(slidingWindowCount)

	if curCount >= l.limit {
		return false
	}

	l.curCount++
	return true
}

func (l *SlidingWindow) Wait() {
	for {
		if l.Allow() {
			break
		}

		// Because we have a sliding window, we will wait for a half of the period and retry
		time.Sleep(l.interval / 2)
	}
}
