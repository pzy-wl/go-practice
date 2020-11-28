/*@Time : 2020/10/21 2:31 下午
@Author : ccc
@File : echo3
@Software: GoLand*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//	使用join包连接字符串
	fmt.Println(strings.Join(os.Args[1:], ""))
}
