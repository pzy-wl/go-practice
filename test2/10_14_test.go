/**
  @author:pzy
  @date:2020/10/14
  @note:
**/
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/vhaoran/vchat/lib/ylog"
)

//Go语言常见坑
func TestBug(t *testing.T) {
	//当参数的可变参数类型是空接口时,传入空接口的切片时需要注意参数的展开问题
	var a = []interface{}{1, 2, 3}
	var b = []interface{}{"2", "3", "4"}
	fmt.Println(a)
	fmt.Println(a...)
	fmt.Println(b)
	fmt.Println(b...)

}

func TestBug2(t *testing.T) {
	//数组是传值,无法通过数组类型的参数修改返回结果
	x := [3]int{1, 2, 3}

	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr)
	}(x)

	fmt.Println(x)
}

func TestBug3(t *testing.T) {
	//map 遍历时顺序不固定
	m := map[string]string{
		"1": "1",
		"2": "2",
		"3": "3",
	}
	for k, v := range m {
		println(k, v)

	}
}

func TestBUg4(t *testing.T) {
	//	在局部作用域中,命名的返回值内同名的局部变量屏蔽
	fmt.Println("---------", test())
}

func test() (n string) {
	n = "hello"
	//return后面省略了n
	return
}

//golang省略返回值是注意内存泄漏问题(省略的返回值是一个指针类型时可能导致)
func TestBug5(t *testing.T) {
	//recover捕获的是祖父级调用时的异常,直接调用时无效
	recover()
	panic(1)
}

func TestBug51(t *testing.T) {
	//直接defer调用也是无效
	defer recover()
	panic(1)
}

func TestBug52(t *testing.T) {
	//defer调用时多层嵌套也是无效
	defer func() {
		func() { recover() }()
	}()
	panic(1)
}

func TestBug53(t *testing.T) {
	//必须是defer函数中直接调用才可以  --生效, 没有panic
	defer func() {
		recover()
	}()
	panic(1)
}

func TestBug6(t *testing.T) {
	//后台goroutine无法保证完成任务
	go println("hello")
}

func TestBug7(t *testing.T) {
	//通过sleep避免并发只中的问题
	for i := 0; i < 2000000; i++ {
		go println("hello", i)
	}
	for i := 0; i < 2000000; i++ {
		go println("hello", i)
	}
	for i := 0; i < 2000000; i++ {
		go println("hello", i)
	}
	for i := 0; i < 2000000; i++ {
		go println("hello", i)
	}
	for i := 0; i < 2000000; i++ {
		go println("hello", i)
	}
	time.Sleep(1 * time.Millisecond) //sleep会切换,等到时间过后不会立即获得cpu而是进入就绪状态,进入等待队列排队
}

func TestBug71(t *testing.T) {
	//或者插入调度语句
	//runtime.GOMAXPROCS(1)
	go println("hello")
	//下面一行代码有让主线程先让出cpu的意思(自己理解的,待完善)
	runtime.Gosched()
	println("hello main")
}

func TestBugn(t *testing.T) {
	// 造场景，设置为单核那么就只能是并发，因为go1.5版本之后，默认是多核了。
	//runtime.GOMAXPROCS(1)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go")
		}
	}()

	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println("hello")
	}
}

func TestBug8(t *testing.T) {
	//独占CPU导致其他Goroutine饿死,Goroutine是协作式抢占调度,自身不会主动放弃cpu
	//	解决办法1:for循环加入 runtime.Gosched()
	//  解决办法2:通过阻塞的方式避免CPU占用 select{}
	runtime.GOMAXPROCS(1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		os.Exit(0)
	}()
	for {
		//println("hello")
		runtime.Gosched()
	} // 占用CPU 1
	//select {} //通过阻塞的方式避免CPU占用
}

//不同GoRoutine之间不满足顺序一致性内存模型
var msg string
var done bool

func setup() {
	msg = "hello, world"
	done = true
}
func TestBug9(t *testing.T) {
	//因为在不同的Goroutine之中,main函数无法保证打印出hello,world:
	go setup()
	for !done {
	}
	println(msg)
}

//解决的办法是显式同步
var msg1 string
var done1 = make(chan bool)

func setup1() {
	msg1 = "hello, world"
	done1 <- true
}

func TestBug91(t *testing.T) {
	go setup1()
	<-done1
	println(msg1)
}

//以下三个测试都有可能造成资源泄漏,因为都在for循环中调用了defer
func TestBug10(t *testing.T) {
	//	闭包错误引用同一变量-- 会在循环结束之后调用defer五次,因此输出结果全是5
	for i := 0; i < 5; i++ {
		defer func() {
			println(i)
		}()
	}
}

func TestBug101(t *testing.T) {
	//解决办法1:是每轮迭代生成一个局部变量
	for i := 0; i < 5; i++ {
		i := i
		defer func() {
			println(i)
		}()
	}
}

func TestBug102(t *testing.T) {
	//	解决办法2:通过函数参数传入
	for i := 0; i < 5; i++ {
		defer func(i int) {
			println(i)
		}(i)
	}
}

func TestDir(t *testing.T) {
	//	获取本地路径
	spath, err := os.Getwd()
	if err != nil {
		return
	}
	f, err := os.Create("2.txt")
	if err != nil {
		ylog.Debug("err is", err)
	}
	spath = spath + "/2.txt"
	fmt.Println(spath)
	_, err = f.WriteString("/Users/ccc/mydata/day3/go-practice/test2/1.txt")
	if err != nil {
		ylog.Debug("文件写入时出错", err)
	}
}

//在循环内部执行defer
func TestBug11(t *testing.T) {
	//	defer在函数退出时才能执行,在for循环执行的defer会导致资源延迟释放:
	//报错(too many open files in system) 的临界值
	for i := 0; i < 7622; i++ {
		f, err := os.Open("/Users/ccc/mydata/day3/go-practice/test2/1.txt")
		if err != nil {
			log.Fatal(err)
		}
		//可能导致资源泄漏,因为defer定义在for循环中了
		defer f.Close()
	}
}

func TestBug111(t *testing.T) {
	//解决办法: 在for中构造一个局部函数,在局部函数内部执行defer
	for i := 0; i < 5; i++ {
		func() {
			f, err := os.Open("/Users/ccc/mydata/day3/go-practice/test2/1.txt")
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
		}()
	}
}

func TestBug12(t *testing.T) {
	//	Goroutine泄漏
	//后台Goroutine向管道输入自然数队列,main函数中输出序列. 但是当break跳出for循环时,后台的Goroutine就处于无法被回收的状态了
	ch := func() <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; ; i++ {
				ch <- i
			}
		}()
		return ch
	}()

	for v := range ch {
		fmt.Println(v)
		if v == 5 {
			break
		}
	}
}

func TestBug121(t *testing.T) {
	//通过context包避免问题
	//原理:当main 函数在break跳出循环时,cancel()来通知Goroutine退出,这样就避免了Goroutine泄漏
	ctx, cancel := context.WithCancel(context.Background())

	ch := func(ctx context.Context) <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; ; i++ {
				select {
				// 如果main函数执行break,则告知此Goroutine
				case <-ctx.Done():
					return
				case ch <- i:
				}
			}
		}()
		return ch
	}(ctx)

	for v := range ch {
		fmt.Println(v)
		if v == 5 {
			//break 之前执行cancel来告知Goroutine
			cancel()
			break
		}
	}
}

func TestFun1(t *testing.T) {
	//	三元表达式 go语言明确不支持

}
