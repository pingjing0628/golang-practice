package main

import (
	"io"
	"log"
	"net"
	"time"
)
// 週期性輸出時間的 TCP 伺服器
// Print out current time each second
func main()  {
	listener, err := net.Listen("tcp", "localhost:7000")

	if err != nil {
		log.Fatal(err)
	}

	for {
		// Accept 會阻斷直到有連線請求進來
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err) // Ex: give up to connect
			continue
		}

		go handleConne(conn) // 並行處理連線
	}
}

// 處理用戶端連線
func handleConne(c net.Conn)  {
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))

		if err != nil {
			return // ex: user failed connect
		}

		time.Sleep(1 * time.Second)
	}
}