package main

import "fmt"

func zeroValue (intVal int) {
	intVal = 0
}

func zeroPointer (intPointer *int) {
	*intPointer = 0
}

func main () {
	i := 1
	fmt.Println("i", i)

	zeroValue(i)
	fmt.Println(i)

	zeroPointer(&i)
	fmt.Println(i)
	fmt.Println(&i)
	fmt.Println(*&i)
}
