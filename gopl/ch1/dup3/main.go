/*@Time : 2020/10/21 4:44 下午
@Author : ccc
@File : main
@Software: GoLand*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//改进dup2  只读文件名,不读取标准输入
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
