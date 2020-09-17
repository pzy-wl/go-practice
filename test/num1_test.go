package main

import (
	"fmt"
	"math"
	"testing"
)

//Golang编程百例
//num1
//描述：用 Golang 实现，将四个数进行排列组合。
//
//题目：有 1、2、3、4 这四个数字，能组成多少个互不相同且无重复数字的三位数？都是多少？
//
//答案：https://haicoder.net/case/golang-hundred-cases/golang-1-1.html
func Test_num1(t *testing.T)  {
	l:=[4]int{1, 2, 3, 4}
	num:=0
	for _, v1:= range l{
		for _, v2:= range l {
			if v1==v2{
				continue
			}
			for _, v3:= range l {
			     if v3==v1||v3==v2{
			     	continue
				 }
				 num++
			}
			}
	}
	println("最终个数是:", num)
}
func Test_num3(t *testing.T)  {
//	完全平方数
//一个整数，它加上 100 后是一个完全平方数，再加上 168 又是一个完全平方数，请问该数是多少？
	i := 0
	for {
		x := int(math.Sqrt(float64(i + 100)))
		y := int(math.Sqrt(float64(i + 268)))

		if x*x == i+100 && y*y == i+268 {
			fmt.Printf("这个数字是 %d\n", i)
			break
		}
		i++
	}
}