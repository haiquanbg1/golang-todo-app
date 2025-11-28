package repositories

type TodoRepository interface {
	Demo() string
}

type todoRepository struct{}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (repo *todoRepository) Demo() string {
	return "demo"
}
