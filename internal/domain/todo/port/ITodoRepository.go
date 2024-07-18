package port

import "github.com/HugoBeaupelFCMS/DevopsJourney/internal/domain/todo/model"

type ITodoRepository interface {
	ReadTodo(int) (model.Todo, error)
	ReadTodoList() ([]model.Todo, error)
	CreateTodo(todo model.Todo) (int, error)
	DeleteTodo(id int) error
	UpdateTodo(todo model.Todo) error
}
