/*@Time : 2020/10/22 11:53 上午
@Author : ccc
@File : main
@Software: GoLand*/
package main

import (
	"fmt"

	"github.com/vhaoran/vchat/common/g"
)

func main() {
	m := make(map[string]interface{})
	m["master"] = 1
	m["mall"] = 123
	m["price"] = 121
	m["vie"] = 134
	for k, v := range m {
		fmt.Println("map内容是", k, v)
		if g.InSlice(k, []string{"master", "mall", "price", "vie"}) {
			fmt.Printf("%s的内容在范围之内\n", k)
		}
	}
}
