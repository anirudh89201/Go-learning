package main

import (
	"fmt"
	"sync"
	"time"
)

func Worker(idx int, jobs chan string, outputChannel chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		time.Sleep(3 * time.Second)
		outputChannel <- fmt.Sprintf("Worker %d is processing %s", idx, job)
	}
}
func main() {
	queue := []string{"A", "B", "C", "D", "E", "F"}
	jobs := make(chan string, len(queue))
	outputChannel := make(chan string, len(queue))
	var wg sync.WaitGroup
	for i := range 3 {
		wg.Add(1)
		go Worker(i, jobs, outputChannel, &wg)
	}
	for _, job := range queue {
		jobs <- job
	}
	close(jobs)
	go func() {
		wg.Wait()
		close(outputChannel)
	}()
	for val := range outputChannel {
		fmt.Println(val)
	}
}
