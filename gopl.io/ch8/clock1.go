package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var flagvar int

// TZ=US/Eastern go run clock1.go -p 3000
// TZ=Asia/Tokyo go run clock1.go -p 3000
func main() {
	flag.IntVar(&flagvar, "p", 3000, "flag port")
	flag.Parse()

	lc := parseLocale()

	listener, e := net.Listen("tcp", "localhost:"+strconv.Itoa(flagvar))
	if e != nil {
		panic(e)
	}
	defer listener.Close()
	for {
		conn, i := listener.Accept()
		if i != nil {
			panic("cannot accept connection")
		}
		go handleConn(conn, lc)
	}
}

func parseLocale() *time.Location {
	// default Asia/Tokyo
	locale := os.Getenv("TZ")
	if locale == "" {
		locale = "Asia/Tokyo"
	}
	lc, err := time.LoadLocation(locale)
	if err != nil {
		log.Fatalf("cannot parse locale: %s \n%v", locale, err)
	}
	return lc
}

func handleConn(conn net.Conn, lc *time.Location) {
	defer conn.Close()
	for {
		io.WriteString(conn, time.Now().In(lc).Format("15:04:05")+"\n")
		time.Sleep(time.Second)
	}
}
