package entity

import "time"

type User struct {
	ID        ID
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
