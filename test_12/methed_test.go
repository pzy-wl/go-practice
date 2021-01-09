/*@Time : 2020/12/8 3:31 下午
@Author : ccc
@File : methed_test
@Software: GoLand*/
package main

import (
	"fmt"
	"image/color"
	"math"
	"net/url"
	"testing"
)

type Point1 struct{ X, Y float64 }

func (p *Point1) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
func (p *Point1) Distance(q Point1) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)

}

// traditional function
// 传统的函数
func Distance(p, q Point1) float64 {
	//math.Hypot是求输入两个参数(float64)的平方和
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point1 type
//同样的事情,但是是作为一个point1类型的方法
//p又称为方法的接收器(receiver)
//func (p Point1) Distance(q Point1) float64 {
//	return math.Hypot(q.X-p.X, q.Y-p.Y)
//}

func TestMethod(t *testing.T) {
	//方法接收器  因为方法和字段都在同一个命名空间, 所以方法与字段不可重复(名称)
	p := Point1{1, 2}
	q := Point1{4, 6}
	fmt.Println(Distance(p, q)) // "5", function call   使用函数
	fmt.Println(p.Distance(q))  // "5", method call     使用p的方法
	fmt.Println(q.Distance(p))  // "5", method call		使用q的方法
}

func Test_main(t *testing.T) {
	//!+main
	// url.Values 是一个键为string,值为字符串数组的map:--->map[string][]string
	m := url.Values{"lang": {"en", "zh"}} // direct construction
	m.Add("item", "1")
	m.Add("item", "2")
	m.Add("item", "3")
	m.Add("item", "4")
	m.Add("lang", "jp")

	//get是获取与给定key相关的第一个值
	fmt.Println(m.Get("lang")) // "en"  如果其对应的值是一个数组,则返回数组中的第一个元素
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // "1"      (first value)
	fmt.Println(m["item"])     // "[1 2]"  (direct map access)
	//
	m = nil
	fmt.Println(m.Get("item")) // ""  对于一个空map查询结果必定为空
	//m.Add("item", "3")         // panic: assignment to entry in nil map  在nil映射中赋值:尝试更新空map
	//!-main
	rgba := color.RGBA{A: 111, B: 111, R: 255, G: 244}
	fmt.Println("颜色值是:", rgba)
}
func TestMethodValue(t *testing.T) {
	//	方法值
	p := Point1{1, 2}
	q := Point1{4, 6}
	dis := p.Distance
	scale := q.ScaleBy
	fmt.Println(dis(q))
	scale(2)
	fmt.Println("scaleBy 1:", q)
	scale(3)
	fmt.Println("scaleBy 3:", q)
}
