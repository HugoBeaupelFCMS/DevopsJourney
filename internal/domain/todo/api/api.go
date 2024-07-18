package api

import (
	"time"

	"github.com/HugoBeaupelFCMS/DevopsJourney/internal/domain/todo/model"
	"github.com/HugoBeaupelFCMS/DevopsJourney/internal/domain/todo/port"
)

type TodoAPI struct {
	todoRepository port.ITodoRepository
}

func NewApi(repo port.ITodoRepository) TodoAPI {
	return TodoAPI{repo}
}

func (api TodoAPI) CreateTodo(req port.ICreationRequest) (int, error) {
	todo := model.Todo{
		ID:           0,
		CreationDate: time.Now(),
		Description:  req.Description(),
		DueDate:      time.Unix(req.DueDate(), 0),
		Title:        req.Title(),
	}
	idTodo, err := api.todoRepository.CreateTodo(todo)
	if err != nil {
		return -1, err
	}
	return idTodo, nil
}

func (api TodoAPI) ReadTodo(id int) (model.Todo, error) {
	return model.Todo{}, nil
}

func (api TodoAPI) ReadTodoList() ([]model.Todo, error) {
	return nil, nil
}

func (api TodoAPI) DeleteTodo(id int) error {
	return nil
}

func (api TodoAPI) UpdateTodo(todo model.Todo) error {
	return nil
}
