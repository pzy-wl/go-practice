package t01

import (
	"fmt"
	"testing"
	"time"

	"github.com/vhaoran/vchat/common/g"
	"github.com/vhaoran/vchat/lib/ylog"
)

// go test -run=Test_int
func Test_int(t *testing.T) {
	a := 0
	b := int64(1)
	c := int8(1)
	d := int16(1)
	e := int32(40)
	fmt.Println("---", b, "---", a, c, d, e)
	fmt.Println("输出测试")
	fmt.Println("123")
}

func Test_string(t *testing.T) {
	fmt.Println("-----------------")
	go func() {
		fmt.Println("hello")
	}()
	go func() {
		fmt.Println("hello 2")
	}()
	go func() {
		fmt.Println("hello  3")
	}()

	time.Sleep(time.Second * 3)
}

func Test_ykit(t *testing.T) {
	b := g.InSlice(1, []int{1, 2, 3, 4})
	fmt.Println("-----------------")
	fmt.Println(b)
	ylog.Debug()

	//ylog.Debug("--------base_test.go--->--b=", b)
	//
	//ylog.DebugDump("--------base_test.go------")
	//ylog.Debug("aaaa")
}
