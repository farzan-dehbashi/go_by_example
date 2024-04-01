package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	should_return := make(chan bool)

	go func() {
		for {
			select {
			case <-should_return:
				return
			case t := <-ticker.C:
				fmt.Println(t)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	ticker.Stop()
	should_return <- true
	fmt.Println("done")
}
