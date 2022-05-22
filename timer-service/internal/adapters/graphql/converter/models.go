package converter

import (
	"strconv"

	"pomodoro.news/timer/internal/adapters/graphql/model"
	"pomodoro.news/timer/internal/domain/entity"
	"pomodoro.news/timer/internal/domain/entity/status"
)

func NewPomodoroModel(e *entity.Pomodoro) *model.Pomodoro {
	return &model.Pomodoro{
		ID:        strconv.Itoa(int(e.ID)),
		Remaining: e.RemainingTime(),
		Status:    NewPomodoroStatus(e),
	}
}

func NewPomodoroStatus(e *entity.Pomodoro) string {
	if e.RemainingDuration() <= 0 {
		return "finished"
	}
	switch e.Status {
	case status.PomodoroStarted:
		return "started"
	case status.PomodoroCancelled:
		return "cancelled"
	case status.PomodoroPaused:
		return "paused"
	case status.PomodoroFinished:
		return "finished"
	default:
		return "N/A"
	}
}
