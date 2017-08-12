package main

import (
	"fmt"
	"reflect"
)

func main() {
	// TODO katsukamaru : インスタンスの型を調べるにはどうしたらいいのか -> reflect.TypeOf(c1)
	c1 := make(map[string]int)
	fmt.Println(reflect.TypeOf(c1)) //map[string]int
	c2 := map[string]int{}
	fmt.Println(reflect.TypeOf(c2)) //map[string]int
}
