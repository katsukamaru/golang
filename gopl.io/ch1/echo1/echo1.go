package main

import (
	"fmt"
	"os"
	"strings"
)

// コマンドライン引数を表示する
func main() {
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	fmt.Println(args)
	fmt.Println(argsWithProgName(os.Args))
	fmt.Println(argsWithIndex(os.Args))
	fmt.Println(argsWithProgWithJoin(os.Args))
}

// 1.1
func argsWithProgName(args []string) string {
	var str string
	var appender = " "
	for _, v := range args {
		str += v + appender
	}
	return str
}

// 1.2
func argsWithIndex(args []string) map[int]string {
	m := map[int]string{}
	for i, v := range args {
		m[i] = v
	}
	return m
}

// 1.3
func argsWithProgWithJoin(args []string) string {
	return strings.Join(args, " ")
}
