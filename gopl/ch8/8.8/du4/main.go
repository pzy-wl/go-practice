/*@Time : 2020/10/20 5:42 下午
@Author : ccc
@File : main
@Software: GoLand*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

//!-1

//隔一段时间将root目录大小计算并显示出来 可以中间间间断
func main() {
	// Determine the initial directories.
	//	确定初始的目录
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//!+2
	// Cancel traversal when input is detected.
	//当检测到输入是取消遍历
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte 读取一个字节
		close(done)
	}()
	//!-2

	// Traverse each root of the file tree in parallel.
	//并行的遍历每一个文件树
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// Print the results periodically.
	//定时的打印结果
	tick := time.Tick(500 * time.Millisecond)
	var nfiles, nbytes int64
loop:
	//!+3
	for {
		select {
		case <-done:
			//取消----收到中断信号,立刻排空管道
			// Drain fileSizes to allow existing goroutines to finish.
			//排空管道用来确保已存在的goroutine结束
			for range fileSizes {
				// Do nothing.
				//只是遍历,不读取也不输出
			}
			return
		case size, ok := <-fileSizes:
			// ...
			//!-3
			//计数
			if !ok {
				break loop // fileSizes was closed 管道关闭
			}
			nfiles++
			nbytes += size
		case <-tick:
			//定时输出进度
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes) // final totals 总数
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
//!+4
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		//如果发出取消指令,直接中断程序
		return
	}
	for _, entry := range dirents(dir) {
		// ...
		//!-4
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
		//!+4
	}
}

//!-4

var sema = make(chan struct{}, 20) // concurrency-limiting counting semaphore  当前限制的信号计数量

// dirents returns the entries of directory dir.
//	返回的是该目录下的键值对
//!+5
func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // acquire token 获取令牌
	case <-done:
		return nil // cancelled 取消
	}
	defer func() { <-sema }() // release token 释放令牌

	// ...read directory...
	//!-5

	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	defer f.Close()

	entries, err := f.Readdir(0) // 0 => no limit; read all entries
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		// Don't return: Readdir may return partial results.
	}
	return entries
}
