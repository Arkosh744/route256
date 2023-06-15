package rate_limiter

import "context"

type RateLimiter interface {
	Wait(ctx context.Context) error
}
