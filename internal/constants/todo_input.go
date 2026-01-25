package constants

type TodoInput struct {
	Task        *string     `json:"task"`
	Description *string     `json:"description"`
	Status      *TodoStatus `json:"status"`
}
