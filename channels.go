package main

import "fmt"

func main() {
	channel := make(chan string)

	go func() { channel <- "message 1" }()

	go func() { channel <- "message 2" }()

	msg := <-channel
	fmt.Println(msg)
	msg = <-channel
	fmt.Println(msg)

}
