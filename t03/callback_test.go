package t03

import (
	"fmt"
	"testing"
)

func Test_CallBack(t *testing.T) {
	sayhello("john", addperfix)
}

func addperfix(perfix, name string) {
	fmt.Println(perfix, "!", name)
}

func sayhello(name string, f func(string, string)) {
	//hello 此处执行回调函数
	f("hello", name)
}
