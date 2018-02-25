package main

import (
	"fmt"
	"time"
)

// this program outputs squared numbers. Like, 0,1,4,9...but will panic.
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
			time.Sleep(time.Second * 1)
		}
		// close(naturals)
		// gopl.io p264
		// チャンネルが閉じられた後にチャンネルに対して送信処理を行うとパニックになる
		// 閉じられた後のチャネルが空になった後、更に受信しようとするとセロ値が返る
		//
		// つまり、naturalsを閉じた後に、更に値を出そうとすると0(intのゼロ値)が返る。
		// 閉じない状態で値を出そうとするとpanicが起こる。
	}()

	go func() {
		for range naturals {
			x := <-naturals
			squares <- x * x
		}
		close(squares)
	}()

	for {
		fmt.Printf("%d\n", <-squares)
	}
}
