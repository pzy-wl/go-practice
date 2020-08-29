package main

import (
	"fmt"
	"testing"
)

//测试控制台的输入输出   只能在main函数里才能使用scanf

func Test_1(t *testing.T)  {
	var i int
	println("请输入一个数字")
	fmt.Scanf("%d", &i)
	println("输入的数字是", i)
}
