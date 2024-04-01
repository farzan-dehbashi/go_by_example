package main

import (
	"fmt"
	"time"
)

func fun(input string) {
	for i := 0; i < 10; i++ {
		fmt.Println(input, ":", i)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	go fun("1")
	go fun("2")

	time.Sleep(time.Second)
	fmt.Println("done")
}
