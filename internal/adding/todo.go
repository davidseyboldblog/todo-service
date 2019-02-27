package adding

type Todo struct {
	UserID      string `json:"userId"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
}

type TodoID struct {
	ID int64 `json:"id"`
}
