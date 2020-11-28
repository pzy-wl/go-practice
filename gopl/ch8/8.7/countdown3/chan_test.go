/*@Time : 2020/10/20 3:56 下午
@Author : ccc
@File : chan_test
@Software: GoLand*/
package main

import (
	"fmt"
	"testing"
)

func TestChan(t *testing.T) {
	ch := make(chan int, 1)
	//ch := make(chan int)
	for i := 0; i < 10; i++ {
		//没循环一次,只能执行一个case,因此一共执行5次传入,5次去除,先传入后取出,交替进行
		select {
		case x := <-ch:
			fmt.Println(x) // "0" "2" "4" "6" "8"
		case ch <- i:
		}
	}
}
