package scheduler

import (
	ctx "context"
	"time"
)

// Scheduler schedules a function in intervals
type Scheduler struct {
	duration time.Duration
	f        func()
}

// Start initiates the function scheduling
func (s Scheduler) Start(ctx ctx.Context) {
	ticker := time.NewTicker(s.duration)
	for {
		select {
		case <-ticker.C:
			s.f()
		case <-ctx.Done():
			return
		}
	}
}

// New creates a new scheduler
func New(duration time.Duration, f func()) Scheduler {
	return Scheduler{
		duration,
		f,
	}
}
