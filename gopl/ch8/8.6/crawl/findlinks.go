/**
  @author:pzy
  @date:2020/10/19
  @note:
**/
package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

//!+crawl
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!-crawl

//!+main
func main() {
	worklist := make(chan []string)

	// Start with the command-line arguments.
	//从命令行参数开始
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.wc -l `find . -name '*.go'`
	// 同时抓取web
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

//!-main

/*
//!+output
$ go build gopl.io/ch8/crawl1
$ ./crawl1 http://gopl.io/
http://gopl.io/
https://golang.org/help/
https://golang.org/doc/
https://golang.org/blog/
...
2015/07/15 18:22:12 Get ...: dial tcp: lookup blog.golang.org: no such host
2015/07/15 18:22:12 Get ...: dial tcp 23.21.222.120:443: socket:
                                                        too many open files
...
//!-output
*/
