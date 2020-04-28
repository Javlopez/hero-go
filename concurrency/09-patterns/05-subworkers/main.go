package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

const (
	WORKERS    = 5
	SUBWORKERS = 3
	TASKS      = 10
	SUBTASKS   = 20
)

func main() {
	var wg sync.WaitGroup

	wg.Add(WORKERS)

	tasks := make(chan int)

	for i := 0; i < WORKERS; i++ {
		go worker(tasks, &wg)
	}

	for i := 0; i < TASKS; i++ {
		tasks <- i
	}

	close(tasks)
	wg.Wait()
}

func worker(tasks chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			return
		}

		subtasks := make(chan int)
		for i := 0; i < SUBWORKERS; i++ {
			go subworker(subtasks)
		}

		fmt.Printf("During the process | Number of CPUS: %d\n", runtime.NumCPU())
		fmt.Printf("During the process | Number of Go routines: %d\n", runtime.NumGoroutine())

		for i := 0; i < SUBTASKS; i++ {
			task1 := task * i
			subtasks <- task1
		}
		close(subtasks)
	}
}

func subworker(subtasks chan int) {
	for {
		task, ok := <-subtasks
		if !ok {
			return
		}
		time.Sleep(time.Duration(task) * time.Millisecond)
		fmt.Printf("[GID:%d, value:%d]\n", getGID(), task)
	}
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
