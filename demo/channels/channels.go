package main

import (
	"fmt"
	"time"
)

type ControlMsg int

type Job struct {
	data   int
	result int
}

const (
	DoExit = iota
	ExitOk
)

func doubler(jobs, results chan Job, control chan ControlMsg) {
	for {
		fmt.Println("Doubler!")
		select {
		case msg := <-control:
			switch msg {
			case DoExit:
				fmt.Println("Exit goroutine")
				control <- ExitOk
				return
			default:
				panic("Unhandled control message")
			}
		case job := <-jobs:
			results <- Job{data: job.data, result: job.data * 2}
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	jobs := make(chan Job, 50)
	results := make(chan Job, 50)
	control := make(chan ControlMsg)

	go doubler(jobs, results, control)

	for i := 0; i < 30; i++ {
		jobs <- Job{i, 0}
	}

	for {
		select {
		case result := <-results:
			fmt.Println(result)
		case <-time.After(50 * time.Millisecond):
			fmt.Println("timed out")
			control <- DoExit
			<-control
			fmt.Println("Program Exit:")
			return
		}
	}
}
