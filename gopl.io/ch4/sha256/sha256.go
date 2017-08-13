package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	a := sha256.Sum256([]byte("x"))
	b := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n",a) // %xはbyte列を16進数に変換して表示する
	fmt.Printf("%x\n",b)
}
