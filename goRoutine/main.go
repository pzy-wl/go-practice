/*@Time : 2021/1/8 9:45 上午
@Author : ccc
@File : main
@Software: GoLand*/
package main

import (
	"flag"
	"fmt"
	"time"
)

//接收的值是一个int指针,默认值是3
var n = flag.Int("n", 3, "number of person")

func main() {
	//斐波那契数列
	//必须有flag.Parse(), 否则接收不到传入的参数,只会使用默认值
	flag.Parse()
	fmt.Println("n is:", *n)
	go spinner(1 * time.Millisecond)
	//const n = 45
	fibN := fib(*n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
