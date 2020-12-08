/*@Time : 2020/12/5 2:43 下午
@Author : ccc
@File : map_test
@Software: GoLand*/
package main

import (
	"fmt"
	"testing"
)

//测试map
func TestMap1(t *testing.T) {
	//测试map的删除   对于map遍历输出时是乱序的
	m := make(map[string]string)
	m["1"] = "hello "
	m["2"] = "world "
	m["3"] = "this "
	m["4"] = "is "
	m["5"] = "pzy "
	delete(m, "2")
	for s, s2 := range m {
		fmt.Println(s, s2)
	}
}
