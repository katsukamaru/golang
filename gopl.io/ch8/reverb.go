package main

import (
	"io"
	"net"
	"time"
)

func main() {
	listener, e := net.Listen("tcp", "localhost:8080")
	if e != nil {
		panic(e)
	}
	defer listener.Close()
	for {
		conn, i := listener.Accept()
		if i != nil {
			panic("cannot accept connection")
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	for {
		io.WriteString(conn, time.Now().Format("15:04:05")+"\n")
		time.Sleep(time.Second)
	}
}
