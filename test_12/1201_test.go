/*@Time : 2020/12/1 11:29 上午
@Author : ccc
@File : 1201_test
@Software: GoLand*/
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func TestPseudoRandom(t *testing.T) {
	// 测试伪随机数  相对随机
	//要得到更随机的数字,就得 加上不同的种子
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: %d\n", i, rand.Intn(2))
	}
}

//golang 压力测试  todo 还没搞懂
func requestServer(ch chan<- string) {
	resp, err := http.Get("http://192.168.1.12:88")
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	ch <- fmt.Sprint(resp.Status)
	defer resp.Body.Close()

}
func TestPressure(t *testing.T) {
	ch := make(chan string)
	count := 5000
	//开始计时
	begin := time.Now()
	fmt.Println("开始时间:", begin)
	for i := 0; i <= count; i++ {
		go requestServer(ch)
	}
	for y := 0; y <= count; y++ {
		fmt.Println(<-ch)
	}
	end := time.Now()
	fmt.Println("结束时间:", end, time.Since(begin))
}

func TestBool1(t *testing.T) {
	//测试布尔值 布尔值可以可&&或||进行结合,并且具有短路行为:
	//如果左边的值已经可以确定整个表达式的值, 运算符右边的值将不会被求值

	//对于空字符串求s[0]容易报空指针异常
	s := ""
	println(s != "" && s[0] == 'x')
	println("panic", s[0])
}

func TestRevers(t *testing.T) {
	//反转
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
