/*@Time : 2020/10/20 3:53 下午
@Author : ccc
@File : main
@Software: GoLand*/
package main

import (
	"fmt"
	"os"
	"time"
)

//带有终止操作的火箭发射倒计时程序
func main() {
	// ...create abort channel...

	//!-

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	//!+
	fmt.Println("Commencing countdown.  Press return to abort.")
	tick := time.NewTicker(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick.C:
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
	//终止goroutine cause the ticker's goroutine to terminate
	tick.Stop()
}

//!-

func launch() {
	fmt.Println("Lift off!")
}
