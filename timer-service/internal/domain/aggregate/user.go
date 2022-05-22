package entity

import (
	"pomodoro.news/timer/internal/domain/entity"
)

type User struct {
	ID       entity.ID
	Username string
}
