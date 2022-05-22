package postgres

import (
	"context"

	"pomodoro.news/timer/internal/adapters/persistence/postgres/dao"

	"github.com/hasanozgan/frodao"
	"pomodoro.news/timer/internal/domain/entity"
	"pomodoro.news/timer/internal/domain/repository"
)

func NewPomodoro() repository.Pomodoro {
	return &pomodoroRepo{
		pomodoroDAO: dao.NewPomodoro(),
	}
}

type pomodoroRepo struct {
	pomodoroDAO *dao.PomodoroDAO
}

func (r *pomodoroRepo) Get(id entity.ID) (*entity.Pomodoro, error) {
	record, err := r.pomodoroDAO.FindByID(context.Background(), frodao.TableIDFromInt(id))
	if err != nil {
		return nil, err
	}
	return record.ToEntity(), nil
}

func (r *pomodoroRepo) GetByUserID(userID entity.ID) ([]*entity.Pomodoro, error) {
	records, err := r.pomodoroDAO.FindByUserID(context.Background(), int(userID))
	if err != nil {
		return nil, err
	}

	results := []*entity.Pomodoro{}
	for _, rec := range records {
		results = append(results, rec.ToEntity())
	}
	return results, nil
}

func (r *pomodoroRepo) Save(e *entity.Pomodoro) error {
	t := dao.NewPomodoroTable(e)

	if t.ID.Get() > 0 {
		return r.pomodoroDAO.Update(context.Background(), t)
	} else {
		record, err := r.pomodoroDAO.Create(context.Background(), t)
		e.ID = entity.ID(record.ID.Get())
		e.Duration = record.Duration
		e.StartedAt = record.StartedAt
		return err
	}
}

func (p *pomodoroRepo) Delete(id entity.ID) error {
	return p.pomodoroDAO.Delete(context.Background(), frodao.TableIDFromInt(int(id)))
}
