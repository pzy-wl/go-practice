/*@Time : 2020/12/26 10:50 上午
@Author : ccc
@File : undressble_test
@Software: GoLand*/
package main

import (
	"fmt"
	"testing"
)

//关于不可寻址的例子

type Named interface {
	// Name 用于获取名字。
	Name() string
}
type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}
func (dog Dog) Name() string {
	return dog.name
}
func Test1(t *testing.T) {
	// 示例1。
	const num = 123
	//num is a const , is can be create but not used
	//_ = &num   // 常量不可寻址。
	//_ = &(123) // 基本类型值的字面量不可寻址。
	println(num)
	num1 := 234
	fmt.Println("num1 is a var, it can't be unused", num1)
}
func Test2(t *testing.T) {
	// 示例2
	var str = "abc"
	_ = str
	//_ = &(str[0]) // 对字符串变量的索引结果值不可寻址。
	//_ = &(str[0:2]) // 对字符串变量的切片结果值不可寻址。
	str2 := str[0]
	_ = &str2 // 但这样的寻址就是合法的。
	//_ = &(123 + 456) // 算术操作的结果值不可寻址。
	num2 := 456
	_ = num2
	//_ = &([3]int{1, 2, 3}[0]) // 对数组字面量的索引结果值不可寻址。
	//_ = &([3]int{1, 2, 3}[0:2]) // 对数组字面量的切片结果值不可寻址。
	_ = &([]int{1, 2, 3}[0]) // 对切片字面量的索引结果值却是可寻址的。
	//_ = &([]int{1, 2, 3}[0:2]) // 对切片字面量的切片结果值不可寻址。
	//_ = &(map[int]string{1: "a"}[0]) // 对字典字面量的索引结果值不可寻址。
}
func Test3(t *testing.T) {
	// 示例3
	var map1 = map[int]string{1: "a", 2: "b", 3: "c"}
	_ = map1
	//_ = &(map1[2]) // 对字典变量的索引结果值不可寻址。
	//_ = &(func(x, y int) int {
	// return x + y
	//}) // 字面量代表的函数不可寻址。
	//_ = &(fmt.Sprintf) // 标识符代表的函数不可寻址。
	//_ = &(fmt.Sprintln("abc")) // 对函数的调用结果值不可寻址。
	dog := Dog{"little pig"}
	_ = dog
	//_ = &(dog.Name) // 标识符代表的函数不可寻址。
	//_ = &(dog.Name()) // 对方法的调用结果值不可寻址。
	//_ = &(Dog{"little pig"}.name) // 结构体字面量的字段不可寻址。
	//_ = &(interface{}(dog)) // 类型转换表达式的结果值不可寻址。
	dogI := interface{}(dog)
	_ = dogI
	//_ = &(dogI.(Named)) // 类型断言表达式的结果值不可寻址。
	named := dogI.(Named)
	_ = named
	//_ = &(named.(Dog)) // 类型断言表达式的结果值不可寻址。
}
func Test4(t *testing.T) {
	// 示例4
	var chan1 = make(chan int, 1)
	chan1 <- 1
	//_ = &(<-chan1) // 接收表达式的结果值不可寻址。
}
func Test5(t *testing.T) {
	//	测试字面量
	println("abc" == "abc") //字面量可以比较
}

func TestMap(t *testing.T) {
	//测试字典
	m := make(map[string]interface{}, 0)
	m["name"] = "张三"
	m["age"] = 18
	m["address"] = []string{"郑州", "南阳", "济南"}

	fmt.Printf("%s的地址是%v", m["name"], m["address"])
}
