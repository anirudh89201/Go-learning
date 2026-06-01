package interfaces

type TaskStore interface {
	Load() ([]Task, error)
	Save(tasks []Task) error
}
