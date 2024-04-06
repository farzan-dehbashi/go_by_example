package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 2)

	channel <- "1"
	channel <- "2"
	fmt.Println(len(channel))

	go func() {
		if len(channel) < 2 {
			channel <- "3"
		}
		time.Sleep(time.Second)
		if len(channel) < 2 {
			channel <- "3"
		}
	}()

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	time.Sleep(time.Second * 2)
	fmt.Println(<-channel)
}
