package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	for i, num := range nums {
		fmt.Println(num)
		fmt.Println(i)
	}
}

// key, value := range dict
