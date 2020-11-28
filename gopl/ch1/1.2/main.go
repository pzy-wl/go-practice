/*@Time : 2020/10/21 2:21 下午
@Author : ccc
@File : main
@Software: GoLand*/
package main

import (
	"fmt"
	"os"
)

func main() {
	//打印命令行参数 并将其拼接成一个字符串打印出来
	var s, sep string
	//for循环遍历命令行参数
	for i := 1; i < len(os.Args); i++ {
		println(i, os.Args[i])
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println("本身的名字是:", os.Args[0])
}
