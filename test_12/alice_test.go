/*@Time : 2020/12/10 2:15 下午
@Author : ccc
@File : alice_test
@Software: GoLand*/
package main

import (
	"fmt"
	"testing"
)

//数组测试
func TestSlice1(t *testing.T) {
	//数组是一种值类型, 而不是指针, 因此可以使用new创建
	//切片复制一定要保证目标切片的容量大于被复制切片,返回的值为实际发生复制的个数
	l := make([]int, 3)
	l[0] = 0
	l[1] = 1
	l[2] = 2
	for _, v := range l {
		fmt.Println(v)
	}
	l1 := make([]int, 5)
	println("实际复制了:", copy(l1, l))
	for k, v := range l1 {
		fmt.Printf("%d:%d\n", k, v)
	}

}
