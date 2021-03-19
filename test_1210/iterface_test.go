/*@Time : 2020/12/10 11:53 上午
@Author : ccc
@File : iterface_test
@Software: GoLand*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

//TODO alt + 鼠标左键可以进行多光标编译
func Test1(t *testing.T) {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	//w = nil
	w.Write([]byte("hello pzy")) //will cause nil pointer dereference
	fmt.Println(w)
	fmt.Println("hello golang")
}
