package repository

import (
	"pomodoro.news/timer/internal/domain/entity"
)

// Pomodoro is a repository
type Pomodoro interface {
	Get(id entity.ID) (*entity.Pomodoro, error)
	GetByUserID(userID entity.ID) ([]*entity.Pomodoro, error)
	Save(entity *entity.Pomodoro) error
	Delete(id entity.ID) error
}
