/**
  @author:pzy
  @date:2020/10/19
  @note:
**/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gopl.io/ch8/thumbnail"
)

//并行是循环迭代的常见并发模型-----生成图片缩略图
func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		thumb, err := thumbnail.ImageFile(input.Text())
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(thumb)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}
