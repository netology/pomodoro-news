package usecase

import (
	"time"

	"pomodoro.news/timer/internal/domain/entity"
	"pomodoro.news/timer/internal/domain/entity/status"
	"pomodoro.news/timer/internal/ports/persistence/repository"
)

const DefaultPomodoroDuration = 25 * 60 // Seconds

func NewPomodoro(repo repository.Pomodoro) Pomodoro {
	return &pomodoroUC{
		repo: repo,
	}
}

type Pomodoro interface {
	Start(userID entity.ID) (*entity.Pomodoro, error)
	Stop(id entity.ID) error
	Pause(id entity.ID) (*entity.Pomodoro, error)
	Continue(id entity.ID) (*entity.Pomodoro, error)
	Get(id entity.ID) (*entity.Pomodoro, error)
	GetByUserID(userID entity.ID) ([]*entity.Pomodoro, error)
}

type pomodoroUC struct {
	repo repository.Pomodoro
}

func (uc *pomodoroUC) Start(userID entity.ID) (*entity.Pomodoro, error) {
	e := &entity.Pomodoro{
		UserID:    userID,
		Duration:  DefaultPomodoroDuration,
		Remaining: DefaultPomodoroDuration,
		StartedAt: time.Now(),
		Status:    status.PomodoroStarted,
	}
	if err := uc.repo.Save(e); err != nil {
		return nil, err
	}
	return uc.repo.Get(e.ID)
}

func (uc *pomodoroUC) Get(id entity.ID) (*entity.Pomodoro, error) {
	return uc.repo.Get(id)
}

func (uc *pomodoroUC) Continue(id entity.ID) (*entity.Pomodoro, error) {
	e, err := uc.repo.Get(id)
	if err != nil {
		return nil, err
	}
	if e.Status != status.PomodoroPaused {
		return e, nil
	}
	e.Status = status.PomodoroStarted
	e.StartedAt = time.Now()
	if err := uc.repo.Save(e); err != nil {
		return nil, err
	}
	return uc.repo.Get(id)
}

func (uc *pomodoroUC) Pause(id entity.ID) (*entity.Pomodoro, error) {
	e, err := uc.repo.Get(id)
	if err != nil {
		return nil, err
	}
	if e.Status == status.PomodoroPaused {
		return e, nil
	}
	e.Status = status.PomodoroPaused
	e.Remaining = e.RemainingDuration()
	if err := uc.repo.Save(e); err != nil {
		return nil, err
	}
	return uc.repo.Get(id)
}

func (p *pomodoroUC) Stop(id entity.ID) error {
	return p.repo.Delete(id)
}

func (uc *pomodoroUC) GetByUserID(userID entity.ID) ([]*entity.Pomodoro, error) {
	return uc.repo.GetByUserID(userID)
}
