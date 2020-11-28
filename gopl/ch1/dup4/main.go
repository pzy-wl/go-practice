/*@Time : 2020/10/21 4:51 下午
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
	//	出现重读行是打印文件的名称 联系1.4
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			if counts[line] > 1 {
				fmt.Println("出现重复行的文件名是:", filename)
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
