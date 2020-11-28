/*@Time : 2020/10/22 9:35 上午
@Author : ccc
@File : main
@Software: GoLand*/
package main

import (
	"fmt"
	"reflect"
)

//检测变量类型
func main() {
	s := "hello, world!"
	n := 123
	p := new(int64)
	a := new(struct{})
	sa := new(string)
	fmt.Println(reflect.TypeOf(s), reflect.TypeOf(n), reflect.TypeOf(p), reflect.TypeOf(a), reflect.TypeOf(sa))
}
