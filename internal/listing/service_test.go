package listing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockRepository struct {
	todo     Todo
	todoList []Todo
	err      error
}

func (m *mockRepository) Get(id int64) (Todo, error) {
	if m.err != nil {
		return Todo{}, m.err
	}
	return m.todo, nil
}

func (m *mockRepository) GetList(userID int64) ([]Todo, error) {
	if m.err != nil {
		return []Todo{}, m.err
	}
	return m.todoList, nil
}

func TestGetTodo(t *testing.T) {

	todo := Todo{
		1,
		1,
		"Description",
		false,
	}

	service := NewService(&mockRepository{
		todo,
		[]Todo{},
		nil,
	})

	actualTodo, _ := service.GetTodo("1")

	assert.Equal(t, todo.ID, actualTodo.ID)
	assert.Equal(t, todo.UserID, actualTodo.UserID)
	assert.Equal(t, todo.Description, actualTodo.Description)
	assert.Equal(t, todo.Complete, actualTodo.Complete)
}

func TestGetTodoList(t *testing.T) {
	todo := Todo{
		1,
		1,
		"Description",
		false,
	}

	service := NewService(&mockRepository{
		Todo{},
		[]Todo{todo},
		nil,
	})

	actualTodoList, _ := service.GetTodoList("1")

	assert.Equal(t, 1, len(actualTodoList.Todos))

	actualTodo := actualTodoList.Todos[0]

	assert.Equal(t, todo.ID, actualTodo.ID)
	assert.Equal(t, todo.UserID, actualTodo.UserID)
	assert.Equal(t, todo.Description, actualTodo.Description)
	assert.Equal(t, todo.Complete, actualTodo.Complete)
}
