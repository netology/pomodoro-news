package graphql

import (
	"pomodoro.news/timer/internal/application/usecase"
)

type Resolver struct {
	PomodoroUC usecase.Pomodoro
}
