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
var num = flag.Int("num", 3, "number of person")

func main() {
	flag.Parse()
	fmt.Println("num is:", *num)
	go spinner(1 * time.Millisecond)
	//const n = 45
	fibN := fib(*num) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", num, fibN)
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
