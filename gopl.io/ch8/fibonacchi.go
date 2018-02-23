package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Fibonacchi culc start.")
	go spinner(time.Millisecond * 100)
	const n = 45
	fib := fib(45)
	fmt.Println("\nFibonacchi culc end.")
	fmt.Printf("Fibonacchi %d is %d", n, fib)

}

func spinner(delay time.Duration) {
	start := time.Now()
	for {
		end := time.Now()
		fmt.Printf("\r%f秒", (end.Sub(start)).Seconds())
		time.Sleep(delay)
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
