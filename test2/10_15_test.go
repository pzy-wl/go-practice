/**
  @author:pzy
  @date:2020/10/15
  @note:
**/
package main

import (
	"fmt"

	"sync"
	"testing"
)

func Inc() (v int) {
	//执行过return之后执行的加以操作
	//此函数称为闭包,闭包对捕获的外部变量并不是传值式访问,而是以引用的方式进行访问
	defer func() { v++ }()
	return 66
}
func TestVar(t *testing.T) {
	//	函数的返回值也可以被命名
	println("返回值是:", Inc())
}

//闭包引发的问题:defer执行引用不准确
func TestDefer(t *testing.T) {
	//三此输出结果全部是3
	for i := 0; i < 3; i++ {
		defer func() { println(i) }()
	}
}

//解决办法: 每轮迭代为的defer函数生成独有的变量
func TestDefer2(t *testing.T) {
	for i := 0; i < 3; i++ {
		// 通过函数传入i
		// defer 语句会马上对调用参数求值
		defer func(i int) {
			println(i)
		}(i)
	}

}

//原子性操作
var total struct {
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 100; i++ {
		//如果不加锁可能导致多个goRoutine竞争total,导致total.value最后不准确
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

func TestSyne(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)

	wg.Wait()

	fmt.Println(total.value)
}
