/*@Time : 2020/10/20 3:28 下午
@Author : ccc
@File : main
@Software: GoLand*/
package main

import (
	"fmt"
	"os"
	"time"
)

//一个带有中断指令的火箭发射倒数模拟程序
func main() {
	abort := make(chan struct{})
	go func() {
		//读取一个字节
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown . Press return to abort")

	select {
	case <-time.After(10 * time.Second):
	//do nothing

	case <-abort:
		//发射倒数完成之前进行其他操作视为终止发射
		fmt.Println("Launch abort!")
		return
	}
	lacuch()
}

func lacuch() {
	fmt.Println("Left off!")
}
