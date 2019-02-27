package listing

import "testing"

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

	if actualTodo.ID != todo.ID {
		t.Fatalf("Expected ID to equal: %v but was: %v", todo.ID, actualTodo.ID)
	} else if actualTodo.UserID != todo.UserID {
		t.Fatalf("Expected UserID to equal: %v but was: %v", todo.UserID, actualTodo.UserID)
	} else if actualTodo.Description != todo.Description {
		t.Fatalf("Expected Description to equal: %v but was: %v", todo.Description, actualTodo.Description)
	} else if actualTodo.Complete != todo.Complete {
		t.Fatalf("Expected Complete to equal: %v but was: %v", todo.Complete, actualTodo.Complete)
	}
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

	if len(actualTodoList.Todos) != 1 {
		t.Fatalf("Expected Len to equal: %v but was: %v", 1, len(actualTodoList.Todos))
	}

	actualTodo := actualTodoList.Todos[0]

	if actualTodo.ID != todo.ID {
		t.Fatalf("Expected ID to equal: %v but was: %v", todo.ID, actualTodo.ID)
	} else if actualTodo.UserID != todo.UserID {
		t.Fatalf("Expected UserID to equal: %v but was: %v", todo.UserID, actualTodo.UserID)
	} else if actualTodo.Description != todo.Description {
		t.Fatalf("Expected Description to equal: %v but was: %v", todo.Description, actualTodo.Description)
	} else if actualTodo.Complete != todo.Complete {
		t.Fatalf("Expected Complete to equal: %v but was: %v", todo.Complete, actualTodo.Complete)
	}
}
