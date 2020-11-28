/*@Time : 2020/10/21 7:09 下午
@Author : ccc
@File : main
@Software: GoLand*/
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	//拼接字符串,第二个参数是拼接时添加上的链接内容
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println("")
	}
}
