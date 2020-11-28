/*@Time : 2020/10/20 4:26 下午
@Author : ccc
@File : main
@Software: GoLand*/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	// Determine the initial directories.
	//确定初始目录
	//从命令行获取初始参数
	flag.Parse()
	roots := flag.Args()
	//返回无标志的命令行参数
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree.
	//遍历文件数
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// Print the results.
	//打印结果
	var nfiles, nbytes int64
	//接收管道传来的结果,并将其相加
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, fileSizes chan<- int64) {
	//递归遍历文件确定文件个数及大小
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents returns the entries of directory dir.
//	返回某个路径下的文件(夹)信息 使用枚举的方法
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
