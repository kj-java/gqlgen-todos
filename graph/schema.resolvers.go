package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gqlgen-todo/gqlgen-todos/entity"
	"github.com/gqlgen-todo/gqlgen-todos/graph/generated"
	"github.com/gqlgen-todo/gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todo := []*model.Todo{
		{
			ID:   "1",
			Text: "サンプルテキスト",
			Done: true,
			User: &model.User{
				ID:   "001",
				Name: "山田太郎",
			},
		},
	}
	return todo, nil
}

func (r *queryResolver) Task(ctx context.Context, id string) (*model.Task, error) {
	idn, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	var t entity.Task
	if err := r.DB.Debug().Find(&t, idn).Error; err != nil {
		return nil, err
	}
	return &model.Task{
		ID:   fmt.Sprintf("%d", t.ID),
		Name: t.Name,
		Task: t.Task,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type taskResolver struct{ *Resolver }
