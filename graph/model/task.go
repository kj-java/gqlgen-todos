package model

import (
	"fmt"
	"github.com/gqlgen-todo/gqlgen-todos/entity"
)

type Task struct {
	ID  string `json:"id"`
	Name string `json:"name"`
	Task string `json:"task"`
}

func NewUserFromEntity(e *entity.Task) *Task {
	return &Task{
		ID:   fmt.Sprintf("%d", e.ID),
		Name: e.Name,
		Task: e.Task,
	}
}