package dao

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/hasanozgan/frodao/drivers/postgres"
	"github.com/hasanozgan/frodao/tableid"
)

func NewPomodoro() *PomodoroDAO {
	return &PomodoroDAO{
		DAO: postgres.NewDAO[PomodoroTable, tableid.Int]("pomodoro"),
	}
}

type PomodoroDAO struct {
	postgres.DAO[PomodoroTable, tableid.Int]
}

func (d *PomodoroDAO) FindByUserID(ctx context.Context, userID int) ([]*PomodoroTable, error) {
	return d.FindByQuery(ctx, d.SelectQuery().Where(goqu.Ex{"user_id": userID}))
}
