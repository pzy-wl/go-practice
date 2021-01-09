/*@Time : 2021/1/8 5:16 下午
@Author : ccc
@File : main
@Software: GoLand*/
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	//无缓存channel
	//一个基于无缓存的channel,发送操作将导致发送者goroutine阻塞,
	//直至另一个goroutine在相同的channels上执行接收操作
	//反之也是,如果接收操作先发生,那么接受者goroutine也将阻塞,
	//知道另一个goroutine在相同的channels上执行发送操作
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // wait for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
