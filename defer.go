package main

import "fmt"

func main() {
	defer func() {
		recover()
		fmt.Println("defer 2")
	}()
	panic("this is a unhandled problem")
}
