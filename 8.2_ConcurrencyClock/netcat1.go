package main

import (
	"io"
	"log"
	"net"
	"os"
)

// 只讀取 TCP 用戶端
func main()  {
	conn, err := net.Dial("tcp", "localhost:7000")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	mustCopy(os.Stdout, conn)
}
// 在不同終端機上同時執行兩個用戶端
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}