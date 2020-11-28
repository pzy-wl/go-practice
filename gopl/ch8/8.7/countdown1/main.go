/*@Time : 2020/10/20 3:21 下午
@Author : ccc
@File : main
@Software: GoLand*/
package main

import (
	"fmt"
	"time"
)

//模拟火箭发射倒计时
func main() {
	fmt.Println("Commencing countdown!")
	tick := time.Tick(1 * time.Second)
	for i := 10; i > 0; i-- {
		fmt.Println("倒计时:", i)
		<-tick
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
