package updating

type Todo struct {
	UserID      string `json:"userId"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
}
