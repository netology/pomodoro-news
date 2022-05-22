package entity

import (
	"time"

	"pomodoro.news/timer/internal/domain/entity/status"
)

type Pomodoro struct {
	ID        ID
	UserID    ID
	Duration  int
	Remaining int
	StartedAt time.Time
	Status    status.Pomodoro
}

func (e Pomodoro) calculateRemaningTime() int {
	total := time.Duration(e.Remaining * int(time.Second))
	now := time.Now()
	diff := now.Sub(e.StartedAt)
	remaining := int((total - diff).Seconds())
	if remaining < 0 {
		return 0
	}
	return remaining
}

func (e Pomodoro) RemainingDuration() int {
	if e.Status == status.PomodoroPaused {
		return e.Remaining
	}
	return e.calculateRemaningTime()
}

func (e Pomodoro) RemainingTime() string {
	out := time.Time{}.Add(time.Duration(e.RemainingDuration() * int(time.Second)))
	return out.Format("04:05")
}
