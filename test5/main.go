/**
  @author:pzy
  @date:2020/10/16
  @note:
**/
package main

import (
	"io"
	"log"
	"net"
	"os"
)

//时钟客户端
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
