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
	// Listen 建構傾聽網路埠連線的 net.Listener 物件
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
		// 回傳代表連線的 net.Conn物件
		handleConn(conn) // Deal with one connection once a time
	}
}

// 處理用戶端連線
func handleConn(c net.Conn)  {
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))

		if err != nil {
			return // ex: user failed connect
		}

		time.Sleep(1 * time.Second)
	}
}