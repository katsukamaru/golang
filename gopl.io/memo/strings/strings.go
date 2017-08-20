package strings

import (
	"fmt"
	"strings"
)

// stringsの関数についてメモをする
func index() {
	// strings.Index は 一番最初に見つけたIndexを返す。見つけなかった場合は-1を返す。
	s := strings.Index("/golang.go", "/")
	fmt.Println(s) // -> 0
	q := strings.Index("golang.go", "/")
	fmt.Println(q) // -> -1
	r := strings.Index("/golang/golang.go", "/")
	fmt.Println(r) // -> 0
}
