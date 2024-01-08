package main

import "fmt"

func main() {
	var s []string
	fmt.Println(s, s == nil, len(s) == 0)
	//	would all be true
	fmt.Println(fmt.Sprintf("%T", s)) // []string

	s = make([]string, 3)
	fmt.Println(s, len(s), cap(s))

	nums := "0123456789"
	fmt.Println(nums[3:])
	//	 no - or reverse traversal

	strings := []string{"11", "22", "33"}
	fmt.Println(strings)
	fmt.Println(strings[0])

	//if slices.Equals(strings, strings) {
	//	fmt.Println("equals")
	//}
}
