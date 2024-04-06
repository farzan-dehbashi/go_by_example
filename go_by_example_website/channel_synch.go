package main

import (
	"fmt"
	"time"
)

func worker(work_in_progress chan bool, name string) {
	for len(work_in_progress) > 0 {
		time.Sleep(time.Millisecond)
	}
	work_in_progress <- true
	fmt.Println(name, "working...")
	time.Sleep(time.Second)
	fmt.Println(name, "done")
	<-work_in_progress
}

func main() {
	work_in_progress := make(chan bool, 1)
	go worker(work_in_progress, "worker 1")
	go worker(work_in_progress, "worker 2")
	time.Sleep(time.Second)
	for len(work_in_progress) > 0 {
		time.Sleep(time.Millisecond)
	}
}
