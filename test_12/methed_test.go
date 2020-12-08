/*@Time : 2020/12/8 3:31 下午
@Author : ccc
@File : methed_test
@Software: GoLand*/
package main

import (
	"fmt"
	"math"
	"testing"
)

type Point1 struct{ X, Y float64 }

// traditional function
// 传统的函数
func Distance(p, q Point1) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point1 type
//同样的事情,但是是作为一个point1类型的方法
func (p Point1) Distance(q Point1) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func TestMethod(t *testing.T) {
	//方法接收器  因为方法和字段都在同一个命名空间, 所以方法与字段不可重复(名称)
	p := Point1{1, 2}
	q := Point1{4, 6}
	fmt.Println(Distance(p, q)) // "5", function call
	fmt.Println(p.Distance(q))  // "5", method call
	fmt.Println(q.Distance(p))  // "5", method call
}
