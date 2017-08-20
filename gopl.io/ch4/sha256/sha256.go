package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	a := sha256.Sum256([]byte("x"))
	b := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n", a) // %xはbyte列を16進 // 数に変換して表示する
	fmt.Printf("%x\n", b)

	count := ByteDiff(a, b)
	fmt.Printf("different byte count is %d", count)
}

// 4.1
func ByteDiff(a [32]byte, b [32]byte) int {
	count := 0
	for i, _ := range a {
		if a[i] != b[i] {
			count++
		}
	}
	return count
}

// 4.2
