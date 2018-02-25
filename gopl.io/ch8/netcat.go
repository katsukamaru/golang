package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {

	conn, e := net.Dial("tcp", "localhost:3000")
	if e != nil {
		panic(e)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
