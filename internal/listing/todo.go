package listing

type Todo struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"userId"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
}

type TodoList struct {
	Todos []Todo `json:"todos"`
}
