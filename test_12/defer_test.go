/*@Time : 2020/12/2 9:36 上午
@Author : ccc
@File : deferTest
@Software: GoLand*/
package main

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	//测试defer 第n次
	println("f的返回值是:", f())
	println("f1的返回值是:", f1())
	println("multiDefer没有返回值,看控制台输出:", multiDefer)
	multiDefer()
}
func f() int {
	//对于未命名的返回值,defer不能改变其值
	//举例:a=10  b=a a++  ----b=10 a不能改变b的值,b相当于返回值
	result := 6
	defer func() {
		// result is accessed after it was set to 6 by the return statement
		result *= 7
	}()
	return result
}

// f returns 42
func f1() (result int) {
	//对于有命名返回值, defer可以改变其值 defer直接对b本身进行操作
	defer func() {
		// result is accessed after it was set to 6 by the return statement
		result *= 7
	}()
	return 6
}
func multiDefer() {
	//对于多个defer函数,顺序逆序执行
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " 前 ")
		}(i)
	}
	println()
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " 后 ")
	}
}

//defer的执行语句为nil, 会产生panic
var fc func() string

func Test2(t *testing.T) {
	//此测试正常结果就是不通过
	println("hello")
	defer fc()
}

func Test3(t *testing.T) {
	//defer传参
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " \n")
		}(i)
		defer func() {
			fmt.Println(i, "")
		}()
	}

}
