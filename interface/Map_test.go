/*@Time : 2021/1/4 3:33 下午
@Author : ccc
@File : Map_test
@Software: GoLand*/
package main

import (
	"fmt"
	"testing"
)

func TestMap1(t *testing.T) {
	m := make(map[string]interface{})
	m["price"] = 1
	m["sort_no"] = 3
	for s, k := range m {
		if s == "price" && k.(int) < 0 {
			fmt.Println("价格不能小于0")
		}
	}
	a, b := m["title"]
	if b {
		fmt.Println("a的值是:", a)
	}
	c, d := m["price"]
	if d {
		fmt.Println("c的值是:", c)
	}
}
