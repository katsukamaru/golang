package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, e := net.Dial("tcp4", "localhost:3000")
	if e != nil {
		panic(e)
	}
	defer conn.Close()
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCp(conn, os.Stdin)
	conn.Close()
	<-done
}

func mustCp(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
