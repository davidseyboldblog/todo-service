package updating

import "testing"

type mockRepository struct {
	err error
}

func (m *mockRepository) Update(id int64, userID int64, desc string, complete int32) error {
	if m.err != nil {
		return m.err
	}
	return nil
}

func TestUpdateTodo(t *testing.T) {
	todo := Todo{
		"1",
		"Description",
		false,
	}

	service := NewService(&mockRepository{
		nil,
	})

	err := service.UpdateTodo("1", todo)

	if err != nil {
		t.Fatal("Expected no error but error was thrown")
	}

}
