/**
  @author:pzy
  @date:2020/10/16
  @note:
**/
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/big"
	"net"
	"testing"
	"time"
	"unsafe"
)

func TestRoutine(t *testing.T) {
	go spinner(100 * time.Millisecond)
	const n = 60
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	//   进入动画函数
	for {
		for _, r := range `-\|/` {
			//println("输出内容之前")
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	//线性递归实现斐波那契数列
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func TestFib(t *testing.T) {
	//	求斐波那契数列
	println(fib(6))
}

func fib2(n int, ret1, ret2 float64) float64 {
	//使用尾递归实现斐波那契数列
	//ret1是最开始第一个数 ret2是第二个数
	if n == 1 {
		return ret1
	}
	//fmt.Printf("%d-------%.2f\n", n, ret1)
	return fib2(n-1, ret2, ret1+ret2)
}

func TestFib2(t *testing.T) {
	//测试尾递归实现的斐波那契数列
	fmt.Printf("%.2f", fib2(1000, 1, 1))
	//println(fib2(1000, 1, 1))
}
func TestInt(t *testing.T) {
	//输出int的取值范围
	fmt.Println("int8:", math.MinInt8, "~", math.MaxInt8)
	fmt.Println("int16:", math.MinInt16, "~", math.MaxInt16)
	fmt.Println("int32:", math.MinInt32, "~", math.MaxInt32)
	fmt.Println("int64:", math.MinInt64, "~", math.MaxInt64)
	//float的取值范围
	//float 的最大值有39位,但是有效值大约是6位,float64的最大值有309位,但是其有效值只有15位
	fmt.Printf("%f---%f", math.MaxFloat32, math.MaxFloat64)
}

func TestFib4(t *testing.T) {
	//	测试数组实现斐波那契数列
	fmt.Printf("%.2f", fib4(100))
}

func fib4(n int) float64 {
	//使用数组实现斐波那契数列
	var l [1000]float64
	l[77] = 5527939700884757
	l[78] = 8944394323791464
	if n < 2 {
		return 1
	}
	for i := 79; i < n; i++ {
		l[i] = l[i-1] + l[i-2]
		fmt.Printf("%d---%.2f\n", i+1, l[i])
	}

	return l[n-1]
}

func TestFloat(t *testing.T) {
	//浮点数骗局  前后不相等
	num := 5527939700884757.0*30 + 8944394323791464.0*30
	fmt.Printf("%f\n", num/30)
	num1 := 5527939700884757 + 8944394323791464
	println("转换前", num1)
	num2 := float64(num1)
	//精确性放大
	fmt.Printf("转换后----%f\n", num2)
	fmt.Printf("转换后加1----%f\n", num2+1.0)
	fmt.Printf("转换后加2----%f\n", num2+2.0)
	fmt.Printf("转换后加3----%f\n", num2+3.0)
	fmt.Printf("转换后加4----%f\n", num2+4.0)
	fmt.Printf("转换后加5----%f\n", num2+5.0)
	fmt.Printf("转换后加6----%f\n", num2+6.0)
}

//测试golang中float的精度
func TestPrecision(t *testing.T) {
	a := 0.1
	b := 0.2
	c := 0.3
	fmt.Printf("%.8f\n", a+b)
	println("结果是", 3.0 == 3)
	if a+b == c {
		fmt.Printf("Isprecision:%v\n", true)
	} else {
		fmt.Printf("Isprecision:%v\n", false)
	}
}
func TestInt2(t *testing.T) {
	//输出数据类型的大小,单位是Byte
	var i1 int = 1
	var i2 int8 = 2
	var i3 int16 = 3
	var i4 int32 = 4
	var i5 int64 = 5
	fmt.Println(unsafe.Sizeof(i1))
	fmt.Println(unsafe.Sizeof(i2))
	fmt.Println(unsafe.Sizeof(i3))
	fmt.Println(unsafe.Sizeof(i4))
	fmt.Println(unsafe.Sizeof(i5))
}

func TestRoutine2(t *testing.T) {
	for {
		for _, r := range `-\|/` {
			//println("输出内容之前")
			fmt.Printf("\r%c", r)
			time.Sleep(5)
		}
	}
}

//并发的clock服务
func TestClock(t *testing.T) {
	listener, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		handleConn(conn) // handle one connection at a time
	}
}

//使用netcat工具执行网络连接操作 --nc localhost 8000
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15-04-05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

func TestFib3(t *testing.T) {
	res := fbn(100)
	fmt.Printf("第个数的斐波那契数是%v", res)
}

//----------------使用big.Int进行斐波那契数列的计算
const LIM = 10000 //求第1000000位的费布拉切数

var fibs [LIM]*big.Int //使用数组保存计算出来的费布拉切数的指针

func TestBigInt(t *testing.T) {
	//使用math.big包进行高位数的斐波那契数的计算
	result := big.NewInt(0)
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		//fmt.Printf("fibonacci(%d) is: %dn", i, result)
	}
	//fibonacci(LIM)
	fmt.Println(result)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %sn", delta)
}

func fibonacci(n int) (res *big.Int) {
	//使用big.Int 进行斐波那契数列高位的运算
	//数组实现 因此必须先求出n前的所有的斐波那契数
	if n <= 1 {
		res = big.NewInt(1)
	} else {
		temp := new(big.Int)
		//两个big.Int 相加
		res = temp.Add(fibs[n-1], fibs[n-2])
	}
	fibs[n] = res
	return
}

func TestBig(t *testing.T) {
	//	测试big.Int
	a := big.NewInt(123)
	fmt.Printf("%v\n", a)

}

//------------------使用big.Int时需要注意的赋值问题:big.NewInt()返回的是一个int类型的指针-----------
func TestBig2(t *testing.T) {

	// 初始化两个变量: a = 1, b = 2
	a := big.NewInt(1)
	b := big.NewInt(2)

	// 打印交换前的数值
	fmt.Printf("a = %v   b = %v\n", a, b)

	// 使用中间变量法进行交换
	tmp := a
	a = b
	b = tmp

	//打印交换后,中间变量不变时的操作结果
	fmt.Printf("中间变量增加前: a = %v    b = %v   tmp = %v\n", a, b, tmp)
	// 交换完成, 对中间变量加100
	tmp.Add(tmp, big.NewInt(100))

	// 打印交换后,中间变量加 100后的结果
	fmt.Printf("中间变量增加后: a = %v    b = %v   tmp = %v\n", a, b, tmp)
}

//-------------------尝试解决,使用赋值而不是引用进行计算--------
func TestBig3(t *testing.T) {
	// 初始化两个变量: a = 1, b = 2
	a := big.NewInt(1)
	b := big.NewInt(2)

	// 打印交换前的数值
	fmt.Printf("a = %v   b = %v\n", a, b)

	// 使用中间变量法进行交换
	tmp := *a
	*a = *b
	*b = tmp

	// 交换完成, 对中间变量加100
	tmp.Add(&tmp, big.NewInt(100))

	// 打印交换后的结果
	fmt.Printf("a = %v    b = %v   tmp = %v\n", a, b, &tmp)
}
