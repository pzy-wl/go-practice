/*@Time : 2020/10/21 2:26 下午
@Author : ccc
@File : echo2
@Software: GoLand*/
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	//for range 遍历命令行参数
	//使用空标识符"_"用于解决语法需要变量名,但是逻辑不需要时的问题
	//定义一个临时变量而不是用会报错!
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
