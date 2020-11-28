/**
  @author:pzy
  @date:2020/10/16
  @note:
**/
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

//Go语言圣经练习8.1 钟表墙
// go run clock.go -p 1234
func main() {
	var port int
	flag.IntVar(&port, "p", 8080, "port")
	flag.Parse()
	fmt.Println(port)

	listener, err := net.Listen("tcp", "localhost:"+strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		_, err := io.WriteString(conn, time.Now().Format(time.RFC3339+"\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
