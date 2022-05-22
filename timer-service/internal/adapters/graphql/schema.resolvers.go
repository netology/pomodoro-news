package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"strconv"

	"pomodoro.news/timer/internal/adapters/graphql/converter"
	"pomodoro.news/timer/internal/adapters/graphql/generated"
	"pomodoro.news/timer/internal/adapters/graphql/model"
	"pomodoro.news/timer/internal/domain/entity"
)

func (r *mutationResolver) Start(ctx context.Context) (*model.Pomodoro, error) {
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		return nil, errors.New("user not found")
	}
	e, err := r.Resolver.PomodoroUC.Start(entity.ID(userID))
	if err != nil {
		return nil, err
	}

	return converter.NewPomodoroModel(e), nil
}

func (r *mutationResolver) Pause(ctx context.Context, id string) (*model.Pomodoro, error) {
	idInt, _ := strconv.Atoi(id)
	if e, err := r.Resolver.PomodoroUC.Pause(entity.ID(idInt)); err != nil {
		return nil, err
	} else {
		return converter.NewPomodoroModel(e), nil
	}
}

func (r *mutationResolver) Stop(ctx context.Context, id string) (bool, error) {
	idInt, _ := strconv.Atoi(id)
	if err := r.Resolver.PomodoroUC.Stop(entity.ID(idInt)); err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) Continue(ctx context.Context, id string) (*model.Pomodoro, error) {
	idInt, _ := strconv.Atoi(id)
	if e, err := r.Resolver.PomodoroUC.Continue(entity.ID(idInt)); err != nil {
		return nil, err
	} else {
		return converter.NewPomodoroModel(e), nil
	}
}

func (r *queryResolver) Pomodoro(ctx context.Context, id string) (*model.Pomodoro, error) {
	idInt, _ := strconv.Atoi(id)
	if e, err := r.Resolver.PomodoroUC.Get(entity.ID(idInt)); err != nil {
		return nil, err
	} else {
		return converter.NewPomodoroModel(e), nil
	}
}

func (r *queryResolver) Pomodoros(ctx context.Context) ([]*model.Pomodoro, error) {
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		return nil, errors.New("not found")
	}

	pomodoros, err := r.Resolver.PomodoroUC.GetByUserID(entity.ID(userID))
	if err != nil {
		return nil, err
	}

	results := []*model.Pomodoro{}
	for _, p := range pomodoros {
		results = append(results, converter.NewPomodoroModel(p))
	}
	return results, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
