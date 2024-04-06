package main

import (
	"fmt"
)

type argError struct {
	arg     int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %d", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 1 {
		return -1, &argError{-1, "cant work with it"}
	}
	return arg * 2, nil
}

func main() {
	_, err = f(1)

}
