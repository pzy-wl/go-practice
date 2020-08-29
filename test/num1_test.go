package main

import (
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