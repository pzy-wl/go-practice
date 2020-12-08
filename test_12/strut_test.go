/*@Time : 2020/12/7 3:11 下午
@Author : ccc
@File : strut_test
@Software: GoLand*/
package main

import (
	"testing"

	"github.com/vhaoran/vchat/lib"
	"github.com/vhaoran/vchat/lib/ylog"
)

func init() {
	_, err := lib.InitModulesOfOptions(&lib.LoadOption{
		LoadMicroService: false,
		LoadEtcd:         false,
		LoadPg:           false,
		LoadRedis:        false,
		LoadMongo:        false,
		LoadMq:           false,
		LoadRabbitMq:     false,
		LoadJwt:          false,
	})
	if err != nil {
		panic(err.Error())
	}
}

//结构体
type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

type Circle1 struct {
	Center Point
	Radius int
}

type Wheel1 struct {
	Circle1 Circle1
	Spokes  int
}

type Circle2 struct {
	Point
	Radius int
}

type Wheel2 struct {
	Circle2
	Spokes int
}

func TestStruct1(t *testing.T) {
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20
	ylog.Debug("输出结果是:", w)
	println(123)
}

func TestStruct2(t *testing.T) {
	//将相同属性独立出来
	//修改后架构提变得很清晰, 但是导致访问每个成员异常繁琐
	var w Wheel1
	w.Circle1.Center.X = 1
	w.Circle1.Center.Y = 1
	w.Circle1.Radius = 2
	w.Spokes = 3
	ylog.Debug(w)
	ylog.DebugDump(w)
}

func TestStruct3(t *testing.T) {
	//引入匿名成员后, 可以直接访问叶子属性而不需要给出完整的路径
	//但是结构体字面值并没有简短表示匿名成员的语法, 因此下面的语句不能编译通过
	//w = Wheel{8, 8, 5, 20}                       // compile error: unknown fields
	//w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} // compile error: unknown fields
	var w Wheel2
	w.Radius = 2
	w.Y = 1
	w.X = 1
	w.Spokes = 3
	println(w.X, w.Y, w.Spokes, w.Radius)
	ylog.DebugDump("输出结果", w)
}

func TestYlog(t *testing.T) {
	//测试ylog为什么失灵了    需要执行init来调用ylog组件
	ylog.Debug("this is a test for ylog")
	println("this is a test for ylog---pl")
}
