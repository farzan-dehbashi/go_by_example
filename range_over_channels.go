package main

import "fmt"

func main() {
	channel := make(chan string, 3)
	channel <- "one"
	channel <- "two"
	close(channel)
	for element := range channel {
		fmt.Println(element)
	}
}
