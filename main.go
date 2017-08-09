package main

import (
	"flag"
	"fmt"
)

var (
	msg string
)

func main() {
	fmt.Println("input message is : " + msg)
}

func init() {
	flag.StringVar(&msg, "str", "blank", "string flag")
	flag.StringVar(&msg, "s", "blank", "string flag")
	flag.Parse()
}
