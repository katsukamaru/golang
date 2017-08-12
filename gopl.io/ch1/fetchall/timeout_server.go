package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Second)
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":8080", nil)
}
