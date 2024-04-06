package main

import "fmt"

func vals() (int, int) {
	return 1, 2
}

func main() {
	fmt.Println(vals())
}
