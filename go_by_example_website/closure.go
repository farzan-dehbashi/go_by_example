package main

import "fmt"

func intCounter() func() int {
    i := 11
    return func() int {
        i--
        return i
    }
}

func main() {
    counter := intCounter()
    fmt.Println(counter())
    fmt.Println(counter())

    new_counter := intCounter()
    fmt.Println(new_counter())
}
