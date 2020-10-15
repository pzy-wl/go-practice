/**
  @author:pzy
  @date:2020/9/28
  @note:
**/
package main

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	//实例1
	var count int
	for {
		select {
		case <-time.Tick(time.Millisecond * 500):
			fmt.Println("咖啡色的羊驼")
			count++
			fmt.Println("count--->", count)
		case <-time.Tick(time.Millisecond * 499):
			fmt.Println(time.Now().Unix())
			count++
			fmt.Println("count--->", count)
		}
	}
}
func TestSelect2(t *testing.T) {
	//实例2
	t2 := time.Tick(500 * time.Millisecond)
	t1 := time.Tick(499 * time.Millisecond)
	var count int
	for {
		select {
		case <-t1:
			fmt.Println("咖啡色的羊驼")
			count++
			fmt.Println("t1 is counting--->", count)
		case <-t2:
			fmt.Println(time.Now().Unix())
			count++
			fmt.Println("t2 is counting--->", count)
		}
	}
}
