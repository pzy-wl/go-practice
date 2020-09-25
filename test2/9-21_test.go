/**
  @author:pzy
  @date:2020/9/21
  @note:
**/
package main

import (
	"fmt"
	"testing"

	. "github.com/vhaoran/vchat/common/ytime"
	"github.com/vhaoran/vchat/lib/ylog"
	. "sync"
)

func TestTime(t *testing.T) {
	ylog.Debug("当前时间是:", Today())
	d := Today()
	fmt.Println(d.Year(), d.Month(), d.Day(), d.Second())
}
func TestOnce(t *testing.T) {
	//sync.Once的Do方法可以实现在程序运行过程中只运行一次其中的回调
	//用到此处虽然多,但是once只执行一次
	var instance int64
	once := Once{}
	once.Do(func() {
		instance = 4
	})
	println(instance)
}
func TestDebug(t *testing.T) {
	//断点测试 测试  F7是step into F8是step over shift+F8是step out
	var a int
	var b int
	var c int
	a = 1
	b = 2
	c = a + b
	fmt.Println(c)
}
