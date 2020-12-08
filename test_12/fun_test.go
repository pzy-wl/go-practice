/*@Time : 2020/12/8 11:17 上午
@Author : ccc
@File : fun_test
@Software: GoLand*/
package main

import (
	"fmt"
	"testing"
)

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

func TestFuc1(t *testing.T) {
	//函数值

	f := square
	fmt.Println(f(3)) // "9"

	f = negative
	fmt.Println(f(3)) // "-3"
	fmt.Printf("%T\n", f)
	f1 := product
	fmt.Println("product(m, n)的值是:", f1(2, 3))
	fmt.Printf("%T\n", f1)
}

func TestFun2(t *testing.T) {
	//函数的空值是nil, 调用会导致panic
	var f func(int) int
	//f(3) //此处f为nil,调用会因此panic错误
	if f != nil {
		f(3)
	}
}

func TestFun3(t *testing.T) {
	//局部匿名变量
	//函数值属于引用类型,且函数值不可比较
	//go用闭包实现函数值,也把函数值成为闭包
	//变量的生命周期不由它的作用于决定,squares返回后,变量x仍然隐式的存在存在于函数值f(f1)中
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
	f1 := squares()
	fmt.Println(f1()) // "1"
	fmt.Println(f1()) // "4"
	fmt.Println(f1()) // "9"
	fmt.Println(f1()) // "16"
}

//调用此函数将生成一个局部变量,并返回一个返回值为该局部变量平方的匿名函数
//生成的匿名函数此函数存在匿名引用
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

//---------------------------------
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
func TestFun4(t *testing.T) {
	//可变参函数
	fmt.Printf("hello, world\n")
	fmt.Println("0:", sum())
	fmt.Println("1:", sum(1))
	fmt.Println("2:", sum(1, 2))
	fmt.Println("3:", sum(1, 2, 3))
	fmt.Println("4:", sum(1, 2, 3, 4))
	values := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(values...))
}

func f2(...int) {}
func g([]int)   {}
func TestFun5(t *testing.T) {
	//可变参数
	fmt.Printf("%T\n", f2) // "func(...int)"
	fmt.Printf("%T\n", g)  // "func([]int)"
}
