package rate_limiter

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"route256/libs/client/pg"
	"sync"
	"time"
)

// SlidingWindowWithPG Do this realisation because store data in pg a bit strange ;).
type SlidingWindowWithPG struct {
	limit    int
	interval time.Duration
	lastTime time.Time `db:"last_time"`

	mu sync.Mutex

	prevCount int `db:"prev_count"`
	curCount  int `db:"cur_count"`

	pg pg.Client
}

const tableName = "rate_limiter_data"

func NewSlidingWindowWithPG(ctx context.Context, limit int, interval time.Duration, pgClient pg.Client) (*SlidingWindowWithPG, error) {
	rl := &SlidingWindowWithPG{
		limit:    limit,
		interval: interval,
		lastTime: time.Now(),
		pg:       pgClient,
	}

	if err := rl.getCount(ctx); err != nil {
		return nil, err
	}

	return rl, nil
}

func (rl *SlidingWindowWithPG) getCount(ctx context.Context) error {
	builder := sq.
		Select("last_time", "prev_count", "cur_count").
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Limit(1)

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := pg.Query{
		Name:     "rate_limiter.getCount",
		QueryRaw: query,
	}

	rlData := struct {
		LastTime  time.Time `db:"last_time"`
		PrevCount int       `db:"prev_count"`
		CurCount  int       `db:"cur_count"`
	}{}

	// this row should exist in db because we create it in migration
	if err = rl.pg.PG().ScanOneContext(ctx, &rlData, q, v...); err != nil {
		return err
	}

	rl.lastTime = rlData.LastTime
	rl.prevCount = rlData.PrevCount
	rl.curCount = rlData.CurCount

	return nil
}

func (rl *SlidingWindowWithPG) setCount(ctx context.Context) error {
	builder := sq.Update(tableName).
		Set("last_time", rl.lastTime).
		Set("prev_count", rl.prevCount).
		Set("cur_count", rl.curCount).
		PlaceholderFormat(sq.Dollar)

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := pg.Query{
		Name:     "rate_limiter.setCount",
		QueryRaw: query,
	}

	if _, err = rl.pg.PG().ExecContext(ctx, q, v...); err != nil {
		return err
	}
	return nil
}

func (rl *SlidingWindowWithPG) Allow(ctx context.Context) (bool, error) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if err := rl.getCount(ctx); err != nil {
		return false, err
	}

	now := time.Now().UTC()
	elapsed := now.Sub(rl.lastTime)

	if elapsed >= rl.interval {
		rl.prevCount = rl.curCount
		rl.curCount = 0
		rl.lastTime = now
	}

	slidingWindowCount := (float64(rl.prevCount) * (rl.interval.Seconds() - elapsed.Seconds())) / rl.interval.Seconds()

	curCount := rl.curCount + int(slidingWindowCount)
	if curCount >= rl.limit {
		rl.curCount = curCount
		if err := rl.setCount(ctx); err != nil {
			return false, err
		}

		return false, nil
	}

	rl.curCount++

	if err := rl.setCount(ctx); err != nil {
		return false, err
	}

	return true, nil
}

func (rl *SlidingWindowWithPG) Wait(ctx context.Context) error {
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
