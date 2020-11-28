/*@Time : 2020/10/21 4:28 下午
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
	//从命令行获取文件名,并且从该文件中读取重复的行以及其重复的次数
	//如果有多个文件,则共用一个map,即两个文件中的重复行也计入统计
	//既可以读取文件名,又可以在没有文件名输入的情况下读取标准输入
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		//如果没有输入参数(文件名)则直接从输入行数中读取并且查找重复行
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				//直接跳过接下来的操作,执行下一次循环
				continue
			}
			countLines(f, counts)
			//关闭该文件释放其资源
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
