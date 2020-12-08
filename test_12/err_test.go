/*@Time : 2020/12/2 3:30 下午
@Author : ccc
@File : err_test
@Software: GoLand*/
package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestErr(t *testing.T) {
	//测试err
	println(fmt.Errorf("%s %v", "错误是:", errors.New("空指针")).Error())
}

func TestBool(t *testing.T) {
	//测试布尔值
	var b bool
	println(b)
}
