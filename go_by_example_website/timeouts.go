package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "done after 2 seconds"
	}()

	select {
	case res := <-channel:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("timed out after 1 second")
	}

	//	now not time out

	go func() {
		time.Sleep(time.Second)
		channel <- "done after 1 second"
	}()

	select {
	case res := <-channel:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timed out")
	}
}
