package adding

import (
	"strconv"

	"github.com/pkg/errors"
)

//Service interface for adding to-dos
type Service interface {
	AddTodo(Todo) (TodoID, error)
}

//Repository used to interact with the data layer.
//This is platform agnostic
type Repository interface {
	Add(int64, string, int32) (int64, error)
}

type service struct {
	todoRepo Repository
}

//NewService returns a service for adding a todo
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddTodo(t Todo) (TodoID, error) {

	//Validation should be done on the input

	userID, err := strconv.ParseInt(t.UserID, 10, 64)
	if err != nil {
		return TodoID{}, errors.Wrap(err, "Error converting id")
	}

	var complete int32
	if t.Complete {
		complete = 1
	} else {
		complete = 0
	}

	todoID, err := s.todoRepo.Add(userID, t.Description, complete)
	if err != nil {
		return TodoID{}, err
	}
	return TodoID{todoID}, nil
}
