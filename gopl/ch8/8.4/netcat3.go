/**
  @author:pzy
  @date:2020/10/19
  @note:
**/
package main

import (
	"io"
	"log"
	"net"
	"os"
)

//测试channel
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors 忽略报错
		log.Println("done")
		done <- struct{}{} // signal the main goroutine 给主线程发送消息
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // wait for background goroutine to finish 等待后台进程结束
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
