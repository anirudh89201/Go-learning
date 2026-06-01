package main

import (
	"fmt"

	"github.com/anirudh89201/hello/01-06-2026/interfaces"
)

func main() {
	store := interfaces.NewJSONStore("tasks.json")
	manager := interfaces.NewManager(store)
	_, err := manager.AddTask("hello World")
	if err != nil {
		fmt.Errorf("Failed to append to tasks.json %w", err)
	}
	tasks, err := manager.GetTasks()
	fmt.Println(tasks)
	for _, val := range tasks {
		fmt.Println(val.ID, val.Description)
	}
}
