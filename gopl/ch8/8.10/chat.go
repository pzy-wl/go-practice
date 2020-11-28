/*@Time : 2020/10/21 9:13 上午
@Author : ccc
@File : chat.go
@Software: GoLand*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//todo pzy
//聊天服务先搁置

type client chan<- string // an outgoing message channel 传出消息通道

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages 所有传入的客户端消息
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients 所有连接的客户端
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			//	对所有客户端传出的消息通道广播即将到来的消息
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

//!-broadcaster
//广播公司
//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages  传出的客户端消息
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()
	//	忽略隐藏的错误
	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors  忽略网络错误
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
