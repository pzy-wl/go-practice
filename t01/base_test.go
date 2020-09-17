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
	println("456")
	println("789")
	println("测试")
	println("测试2")
	println("测试3")
}

func Test_string(t *testing.T) {
	fmt.Println("-----------------")
	//t0:=time.Now()
	//go func() {
	//	fmt.Println("hello")
	//}()
	//go func() {
	//	fmt.Println("hello 2")
	//}()
	//go func() {
	//	fmt.Println("hello  3")
	//}()
	//fmt.Println(time.Since(t0))
	//
	////time.Sleep(time.Second * 3)
	t0 := time.Now()
	for i := 0; i < 10000; i++ {
		go func(i int) {
			fmt.Println("a", i)
		}(i)
	}
	for i := 0; i < 10000; i++ {
		go func(i int) {
			fmt.Println("b", i)
		}(i)
	}
	for i := 0; i < 10000; i++ {
		go func(i int) {
			fmt.Println("c", i)
		}(i)
	}
	for i := 0; i < 10000; i++ {
		go func(i int) {
			fmt.Println("d", i)
		}(i)
	}
	for i := 0; i < 10000; i++ {
		go func(i int) {
			fmt.Println("e", i)
		}(i)
	}
	fmt.Print("此次测试用时:\n", time.Since(t0))

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
