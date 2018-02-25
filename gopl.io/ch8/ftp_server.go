package main

import (
	"log"
	"net"
)

func main() {
	listener, e := net.Listen("tcp", "localhost:1000")
	defer listener.Close()
	if e != nil {
		log.Fatalf("%v", e)
	}

	for {
		conn, e := listener.Accept()
		defer conn.Close()
		if e != nil {
			log.Fatalf("%v", e)
		}
		// 受け取ったコマンドに寄って呼び出す振る舞いを変える
	}
}
