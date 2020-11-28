/*@Time : 2020/10/21 2:39 下午
@Author : ccc
@File : main
@Software: GoLand*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//使用map统计行数,如果相同的行数则同一个key, value+1
	counts := make(map[string]int)
	//ctl+d结束输入
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	//	忽略来自于输入的潜在错误
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
