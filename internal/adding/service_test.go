package adding

import "testing"

type mockRepository struct {
	id  int64
	err error
}

func (m *mockRepository) Add(userID int64, desc string, complete int32) (int64, error) {
	if m.err != nil {
		return -1, m.err
	}
	return m.id, nil
}

func TestAddTodo(t *testing.T) {
	todo := Todo{
		"1",
		"Description",
		false,
	}

	service := NewService(&mockRepository{
		1,
		nil,
	})

	actualTodoID, _ := service.AddTodo(todo)
	if actualTodoID.ID != 1 {
		t.Fatalf("Expected ID to equal: %v but was: %v", 1, actualTodoID.ID)
	}

}
