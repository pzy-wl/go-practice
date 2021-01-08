/*@Time : 2021/1/8 10:50 上午
@Author : ccc
@File : main
@Software: GoLand*/
package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

var n = flag.Int("port", 8000, "listening port")

func main() {
	//新建一个监听tcp的8000端口的接收器
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		//等待下一个连接到接收器
		//Accept waits for and returns the next connection to the listener
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted  连接失败
			continue
		}
		handleConn(conn) // handle one connection at a time  一次处理一个连接(不适用go)
		//	加上go,可以同时处理多个连接
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format(" 15:04:05\n"))
		//_, err := io.WriteString(c, time.Now().String())
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
