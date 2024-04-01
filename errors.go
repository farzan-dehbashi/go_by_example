package main

import (
	"errors"
	"fmt"
)

func function(arg int) (int, error) {
	if arg == 1 {
		return -1, errors.New("cannot work with 1")
	}

	return arg - 1, nil
}

func main() {
	fmt.Println(function(1))
}
