/*@Time : 2020/12/10 11:14 上午
@Author : ccc
@File : main
@Software: GoLand*/
package main

import (
	"flag"
	"fmt"
	"time"

	"gopl.io/ch7/tempconv"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

//var name = flag.Set("name", "zs")
//value是默认值,name是执行时传入的参数-v, 接收的值是一个string
var name = flag.String("name", "zs", "as name")

//value是默认值,enabled执行时传入的参数,接收的值是一个bool指针 只要传入参数,指针对应的值就是true
var enabled = flag.Bool("enabled", false, "module enabled")

//接收的值是一个int指针,默认值是3
var num = flag.Int("num", 3, "number of person")

//默认摄氏温度是20°C, 接收值是一个温度指针,指针对应的是相应的温度值
var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...\n", *period)
	time.Sleep(*period)
	fmt.Println("name is:", *name)
	fmt.Println("module status is:", *enabled)
	fmt.Println("number of person is :", *num)
	fmt.Println("The current temperature is :", *temp)
}
