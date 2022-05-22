package dao

import (
	"time"

	"github.com/hasanozgan/frodao"
	"github.com/hasanozgan/frodao/tableid"
	"pomodoro.news/timer/internal/domain/entity"
	"pomodoro.news/timer/internal/domain/entity/status"
)

type PomodoroTable struct {
	frodao.Table[tableid.Int]

	UserID            int       `db:"user_id"`
	Duration          int       `db:"duration"`
	RemainingDuration int       `db:"remaining_duration"`
	StartedAt         time.Time `db:"started_at"`
	Status            int       `db:"status"`
}

func (t PomodoroTable) ToEntity() *entity.Pomodoro {
	return &entity.Pomodoro{
		ID:        entity.ID(t.Table.ID.Get()),
		UserID:    entity.ID(t.UserID),
		Duration:  t.Duration,
		Remaining: t.RemainingDuration,
		Status:    status.Pomodoro(t.Status),
		StartedAt: t.StartedAt,
	}
}

func NewPomodoroTable(e *entity.Pomodoro) *PomodoroTable {
	return &PomodoroTable{
		Table: frodao.Table[tableid.Int]{
			ID: frodao.TableIDFromInt(int(e.ID)),
		},
		UserID:            int(e.UserID),
		Duration:          e.Duration,
		RemainingDuration: e.Remaining,
		StartedAt:         e.StartedAt,
		Status:            int(e.Status),
	}

}
