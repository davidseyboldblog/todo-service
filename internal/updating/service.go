package updating

import (
	"strconv"

	"github.com/pkg/errors"
)

type Service interface {
	UpdateTodo(string, Todo) error
}

type Repository interface {
	Update(int64, int64, string, int32) error
}

type service struct {
	todoRepo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) UpdateTodo(id string, t Todo) error {
	// Input should be validated

	longID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errors.Wrap(err, "Error converting id")
	}

	userID, err := strconv.ParseInt(t.UserID, 10, 64)
	if err != nil {
		return errors.Wrap(err, "Error converting id")
	}

	var complete int32
	if t.Complete {
		complete = 1
	} else {
		complete = 0
	}

	err = s.todoRepo.Update(longID, userID, t.Description, complete)
	if err != nil {
		return errors.Wrap(err, "Error updating todo")
	}

	return nil
}
