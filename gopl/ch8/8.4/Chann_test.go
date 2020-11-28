/**
  @author:pzy
  @date:2020/10/19
  @note:
**/
package main

import (
	"fmt"
	"testing"
	"time"
)

//测试channel
func TestPipeline(t *testing.T) {
	//定义两个channel
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	// 计数
	go func() {
		//向naturals传递自然数
		for x := 0; ; x++ {
			naturals <- x
			time.Sleep(100 * time.Millisecond)
			if x == 10 {
				break
			}
		}
		close(naturals)
		println("发送端已经关闭")
	}()

	// Squarer
	// 求平方
	go func() {
		for {
			//一直等待管道传出数据
			x, b := <-naturals
			println(b)
			if !b {
				println("退出循环已经执行!")
				break // channel was closed and drained 管道关闭并且枯竭
			}
			squares <- x * x
		}
		close(squares)
		println("接收端已经关闭")
	}()

	// Printer (in main goroutine)
	//在主线程里进行输出 自然数的平方
	for {
		r, b := <-squares
		if b {
			fmt.Println(r)
		} else {
			break
		}
	}
	println("输出端已关闭!")
}

func TestChan2(t *testing.T) {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}

//测试break-----------直接跳出for循环
func TestBreak(t *testing.T) {
	s := make(chan int)
	go func() {
		for i := 0; ; i++ {
			println(i)
			s <- i
			if i == 10 {
				break
			}
		}

	}()
	go func() {
		for {
			f := <-s
			println("接收到的数字是", f)
			if f == 3 {
				break
			}
		}

	}()

}

//封装成函数
func TestChan3(t *testing.T) {
	natural := make(chan int)
	squarer := make(chan int)
	go Count(natural)
	go Squarer(squarer, natural)
	Print(squarer)
}

func Count(out chan int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func Squarer(out, in chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func Print(in chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func TestChan5(t *testing.T) {
	//	向无缓存chan一直发送 无缓存只能值发送一个接收一个,如果没有接收,则无法发送
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()
	//
	//println(<-c)
	//println(cap(c))
}

//------------------带有缓存的channel

func TestChan6(t *testing.T) {
	//	带有缓存的channel 可以没有接收端的前提下预存cap容量的消息
	c := make(chan int, 3)
	go func() {
		for i := 0; i < 4; i++ {
			c <- i
			println(i)
		}
	}()
	//println(<-c)
}
