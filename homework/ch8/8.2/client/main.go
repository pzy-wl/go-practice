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
	"strings"
	"time"
)

//go语言圣经联系题8.1 种表墙
// go run clockwall.go tokyo=1234
func main() {
	for _, v := range os.Args[1:] {
		params := strings.Split(v, "=")
		uri := params[1]
		go connectService(uri)
	}
	for {
		time.Sleep(1 * time.Second)
	}
}

func connectService(uri string) {
	conn, err := net.Dial("tcp", uri)
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
