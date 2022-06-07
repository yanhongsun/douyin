package cache

import (
	"context"
	"time"
)

// Task runs a task at fixed duration which will call fn().
type Task struct {
	// Before will be called before running this task.
	Before func(ctx context.Context)

	// Fn is main function which will called in loop.
	Fn func(ctx context.Context)

	// After will be called after the task loop.
	After func(ctx context.Context)
}

// Run runs this task at fixed duration d.
func (tt *Task) Run(ctx context.Context, d time.Duration) {
	if tt.Fn == nil {
		return
	}

	if tt.Before != nil {
		tt.Before(ctx)
	}

	if tt.After != nil {
		defer tt.After(ctx)
	}

	ticker := time.NewTicker(d)
	defer ticker.Stop()

	for {
		select {
		// Done：返回一个Channel，用于向当前协程传递是否结束；
		case <-ctx.Done():
			return
		case <-ticker.C:
			tt.Fn(ctx)
		}
	}
}
