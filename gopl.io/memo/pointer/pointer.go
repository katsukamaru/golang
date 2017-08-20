package main

import (
	"fmt"
	"reflect"
)

func main() {

	m := map[string]int{}
	fmt.Println(m)
	fmt.Println(reflect.TypeOf(m))

	n := make(map[string]int)
	fmt.Println(n)
	fmt.Println(reflect.TypeOf(n))

	o := new(map[string]int) // new はポインタをかえす。
	fmt.Println(o)
	fmt.Println(reflect.TypeOf(o))

	// 慣習的にポインタ型に対する変数名に決まりとかある？同じ名前だとかなりわかりづらい。
	// map[]
	// map[string]int
	// map[]
	// map[string]int
	// &map[]
	// *map[string]int
	fmt.Println(&o)
	fmt.Println(*o)
	fmt.Println(*o == nil)
	fmt.Println(o["aaa"])

}
