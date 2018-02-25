package main

import (
	"fmt"
	"time"
)

// this program outputs squared numbers. Like, 0,1,4,9...0,0,0,0,0
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
			time.Sleep(time.Second * 1)
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for {
		fmt.Printf("%d\n", <-squares)
	}
}
