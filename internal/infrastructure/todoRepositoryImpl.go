package infrastructure

import (
	"github.com/HugoBeaupelFCMS/DevopsJourney/internal/domain/todo/model"
	"github.com/HugoBeaupelFCMS/DevopsJourney/internal/domain/todo/port"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(URL string, drop bool) (port.ITodoRepository, error) {
	rep := todoRepository{}
	err := rep.InitDatabase(URL, drop)
	return &rep, err
}

func (r *todoRepository) InitDatabase(URL string, drop bool) error {
	var err error

	dialect := sqlite.Open(URL)

	r.db, err = gorm.Open(dialect)
	return err
}

func (r *todoRepository) CreateTodo(todo model.Todo) (int, error) {
	return 0, nil
}

func (r *todoRepository) ReadTodo(id int) (model.Todo, error) {
	return model.Todo{}, nil
}

func (r *todoRepository) ReadTodoList() ([]model.Todo, error) {
	return mapToSummaryResponse(), nil
}

func mapToSummaryResponse() []model.Todo {
	return nil
}

func (r *todoRepository) DeleteTodo(id int) error {
	return nil
}

func (r *todoRepository) UpdateTodo(todo model.Todo) error {
	return nil
}
