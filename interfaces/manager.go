package interfaces

import (
	"fmt"
	"time"
)

type Manager struct {
	taskstore TaskStore
}

func NewManager(store TaskStore) *Manager {
	return &Manager{taskstore: store}
}

func (m *Manager) AddTask(description string) (Task, error) {
	tasks, err := m.taskstore.Load()
	if err != nil {
		return Task{}, fmt.Errorf("Failed to load the task:%w", err)
	}
	newId := 1
	if len(tasks) > 0 {
		newId = tasks[len(tasks)-1].ID + 1
	}
	newTask := Task{
		ID:          newId,
		Description: description,
		Status:      StatusStarted,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now().Add(-3 * time.Minute),
	}
	tasks = append(tasks, newTask)
	if err := m.taskstore.Save(tasks); err != nil {
		return Task{}, fmt.Errorf("Unable to save the task:%w", err)
	}
	return newTask, nil
}
func (m *Manager) GetTasks() ([]Task, error) {
	data, err := m.taskstore.Load()
	if err != nil {
		return nil, fmt.Errorf("Error in loading the data:%w", err)
	}
	return data, nil
}
