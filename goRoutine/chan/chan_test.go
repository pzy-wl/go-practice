/*@Time : 2021/1/8 5:17 下午
@Author : ccc
@File : chan_test
@Software: GoLand*/
package main

import (
	"fmt"
	"testing"
	"time"
)

func TestChan1(t *testing.T) {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}

func TestChan2(t *testing.T) {
	//第一个的改进版
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

//###################################################

func counter(out chan<- int) {
	//参数是一个只发送int不接收的channel
	//生成自然数并通过channel传递
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

//################单向CHANNEL######################

func squarer(out chan<- int, in <-chan int) {
	//两个参数,out 是只发送不接受的channel,in是只接受不发送的channel
	//接收channel发送来的自然数,并且平方后传递给channel out
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	//参数是一个只接受不发送的channel
	//接收并且打印自然数
	for v := range in {
		fmt.Println(v)
	}
}

func TestChan3(t *testing.T) {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	//只能一个发送, 一个接收,否则,前面的撑死, 后面的饿死
	//printer(naturals)
	go squarer(squares, naturals)
	printer(squares)

}

func TestChan4(t *testing.T) {
	//ch <- x  // a send statement  一个发送声明
	//x = <-ch // a receive expression in an assignment statement  赋值语句中的接收表达式
	//<-ch     // a receive statement; result is discarded   一个接收表达式,结果被丢弃
	//单向channel
	//	类型chan<- int, 表示一个只发送int的channel,只能发送不能接收, 类型<-chan int表示一个只接受int的channel,只能接收不能发送
	//双向channel chan int 转换为chan <-int 或者<-chan int 时是隐式转换,但是单向channel不能转换为双向channel
}

//####################带缓存的channel###################

func TestChan5(t *testing.T) {
	//带缓存的channel内部保持了一个元素队列,队列的最大容量, 调用make函数时通过第二个参数指定的
	//向带有缓存的channel执行发送操作,就是往元素队列队尾插入元素,接收操作是从队首删除元素
	//如果内部存储队列是满的,那么发送操作将阻塞,直接列外的goroutine执行接收操作
	ch := make(chan string, 3)
	ch <- "A"
	ch <- "B"
	ch <- "C"
	//内置函数cap可以获取内部缓存队列的容量
	//内置函数len可以获取内部缓存队列的长度
	fmt.Println(len(ch))
	//接收一个参数,并丢弃
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	go func() {
		//延迟发送, 使main goroutine阻塞
		time.Sleep(3 * time.Second)
		ch <- "D"
	}()
	//缓存队列为空后,继续执行接收操作将发生阻塞
	fmt.Println(<-ch)
	fmt.Println(len(ch))
}
