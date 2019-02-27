package listing

import (
	"strconv"

	"github.com/pkg/errors"
)

type Service interface {
	GetTodo(string) (Todo, error)
	GetTodoList(string) (TodoList, error)
}

type Repository interface {
	Get(int64) (Todo, error)
	GetList(int64) ([]Todo, error)
}

type service struct {
	todoRep Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetTodo(id string) (Todo, error) {
	// Validation should be performed on the input

	longID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return Todo{}, errors.Wrap(err, "Error converting id")
	}

	todo, err := s.todoRep.Get(longID)
	if err != nil {
		return Todo{}, errors.Wrap(err, "Error retrieving todo")
	}
	return todo, nil
}

func (s *service) GetTodoList(userID string) (TodoList, error) {
	longID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return TodoList{}, errors.Wrap(err, "Error converting id")
	}

	todoList, err := s.todoRep.GetList(longID)
	if err != nil {
		return TodoList{}, errors.Wrap(err, "Error getting Todo list")
	}
	return TodoList{todoList}, nil
}
