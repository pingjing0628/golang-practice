package server

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var port int

// Package flag implements command-line flag parsing.
// Bind the flag to a variable using the IntVar() functions.
func init()  {
	flag.IntVar(&port, "port", 7000, "port number")
}

func main()  {
	flag.Parse()

	server := fmt.Sprintf("localhost:%d", port)

	listener, err := net.Listen("tcp", server)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Listening at localhost:%d\n", port)

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err) // Ex: connection aborted
			continue
		}

		go handleConn(conn) // handle connections concurrently
	}
}

// 處理用戶端連線
func handleConn(c net.Conn)  {
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))

		if err != nil {
			return // ex: client disconnected
		}

		time.Sleep(1 * time.Second)
	}
}