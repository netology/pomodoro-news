package repository

import (
	"pomodoro.news/timer/internal/domain/entity"
)

// User is a repository
type User interface {
	Get(id entity.ID) (*entity.User, error)
}
