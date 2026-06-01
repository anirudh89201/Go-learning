package interfaces

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonStore struct {
	filePath string
}

func NewJSONStore(filePath string) *JsonStore {
	return &JsonStore{filePath: filePath}
}

func (s *JsonStore) Load() ([]Task, error) {
	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		return []Task{}, nil
	}
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read file %w", err)
	}
	var tasks []Task
	if err = json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("failed to parse json %w", err)
	}
	return tasks, nil
}
func (s *JsonStore) Save(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return fmt.Errorf("Failed to marshal:%w", err)
	}
	err = os.WriteFile(s.filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("Failed to write into file:%w", err)
	}
	return nil
}
