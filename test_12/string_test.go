/*@Time : 2020/12/7 11:46 上午
@Author : ccc
@File : string_test
@Software: GoLand*/
package main

import (
	"strings"
	"testing"
)

func TestToLower(t *testing.T) {
	//测试strings中的tolower
	s := "你好,世界"
	s1 := "HELLO world, ZZZ"
	println(strings.ToLower(s), strings.ToUpper(s), strings.ToLower(s1), strings.ToUpper(s1))
	println(strings.ToUpper(s) == strings.ToLower(s))
	println(strings.Trim(s1, "world"))
}
