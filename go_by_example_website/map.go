package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["key1"] = 1
	m["key2"] = 2
	fmt.Println(m)

	delete(m, "key1")
	val, key_exists := m["k2"]
	fmt.Println(key_exists)
	fmt.Println(val)
	fmt.Println(m)
}
