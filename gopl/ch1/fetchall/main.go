/*@Time : 2020/10/21 5:02 下午
@Author : ccc
@File : main
@Software: GoLand*/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		go fetch(url, ch) // start a goroutine 开始一个goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch 从通道获取一个信号
	}
	//输出一共耗时多久
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch 对通道ch发送错误信息
		return
	}
	fmt.Printf("%s的状态码%s\n", url, resp.Status)
	//ioutil.Discard相当于一个垃圾桶,我们只想要字节数而不想要内容,因此使用ioutil.Discard来过渡一下
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
