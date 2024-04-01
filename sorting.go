package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"a", "d", "b"}
	slices.Sort(strs)
	fmt.Println(strs)
}
