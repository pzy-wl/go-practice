/**
  @author:pzy
  @date:2020/9/25
  @note:
**/
package main

import (
	"testing"
)

func Test_del(t *testing.T) {
	//删除数组中某个元素
	a := []int64{1, 2, 3, 4, 5}
	a = append(a[:2], a[3:]...)
	for _, v := range a {
		println(v)
	}
}
