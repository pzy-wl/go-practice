/*@Time : 2020/11/28 10:48 上午
@Author : ccc
@File : 11-28_test
@Software: GoLand*/
package main

import (
	"fmt"
	"testing"
)

// 错误的调用示例
func TestDefer1(t *testing.T) {
	defer func() {
		doRecover()
	}()
	panic("not good")
}

func doRecover() {
	fmt.Println("recobered: ", recover())
}
