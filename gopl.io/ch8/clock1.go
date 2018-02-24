package main

import (
	"io"
	"net"
	"time"
)

func main() {
	listener, e := net.Listen("tcp", "localhost:3000")
	if e != nil {
		panic(e)
	}
	handleConn(listener)
}

func handleConn(listener net.Listener) {
	defer listener.Close()
	for {
		conn, i := listener.Accept()
		if i != nil {
			panic("cannot accept connection")
		}
		defer conn.Close()
		for {
			io.WriteString(conn, time.Now().Format("15:04:05")+"\n")
			time.Sleep(time.Second)
		}
	}
}
