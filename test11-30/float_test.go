/*@Time : 2020/11/30 2:40 下午
@Author : ccc
@File : float_test
@Software: GoLand*/
package main

import (
	"fmt"
	"math"
	"testing"
)

//测试浮点数
func Test1(t *testing.T) {
	//浮点数的表示 NAN非数一般用于表示无效的除法操作结果,例如:0/0或sqrt(-1)
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)
	//	0 -0 +Inf -Inf NaN
}

func Test2(t *testing.T) {
	//NAN数的不唯一性 在浮点数中,NAN,正无穷大和负无穷小都是不唯一的, 每个都有非常多的bit模式表示
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)
}

func Test3(t *testing.T) {
	//大数吃小数  精度丢失-----float64 的精度大约是15位,float32 的精度大约是6位
	a := 10000000000000000.0
	b := 1.0
	fmt.Printf("%1.30f", a+b)
}

func Test4(t *testing.T) {
	//浮点数精度丢失问题 精度到达小数点后15位左右
	a := 9.1
	fmt.Sprint(a)
	fmt.Printf("%f\n", a)
	fmt.Printf("%.10f\n", a)
	fmt.Printf("%.16f\n", a)
	//	9.100000
	//9.1000000000
	//9.0999999999999996
}

func Test5(t *testing.T) {
	//再次证明浮点数的精度(float64)是小数点后15位
	a := 0.3
	b := 0.6
	fmt.Sprint(a + b)
	fmt.Printf("%f\n", a+b)
	fmt.Printf("%.10f\n", a+b)
	fmt.Printf("%.15f\n", a+b)
	fmt.Printf("%.16f\n", a+b)
	fmt.Printf("%.17f\n", a+b)

}
func Test6(t *testing.T) {
	//float32的精度   表示数值大于16777216就不准确了
	sum := float32(0)
	for i := 0; i < 16777216; i++ {
		a := float32(1.0)
		sum += a
	}
	fmt.Printf("%.6f", sum)
}

func Test7(t *testing.T) {
	//浮点数精度丢失解决办法  算出损失的精度,并在下次计算是追加上  Kahan Summation 算法
	sum := float32(0)
	c := float32(0)
	for i := 0; i < 20000000; i++ {
		x := float32(1.0)
		y := x - c
		t := sum + y
		c = (t - sum) - y
		fmt.Printf("c:%f\n", c)
		sum = t
	}
	fmt.Printf("%.6f\n", sum)
	fmt.Println("结果是", sum)
	//	20000000.000000
	//结果是 2e+07
}

func Test8(t *testing.T) {
	//不能拿浮点数和0比较   当某个浮点数的绝对值小于某个跟小的数(比如十万分之一)就可以认为其值为0
	a := 0.3

	fmt.Println((a - a + a) == a)
	fmt.Printf("%.17f", a)
}

func Test9(t *testing.T) {
	//浮点数精度 0后面323位后加才可能等于0(float64),而float只能紧缺到小数点后45位--------等于0时
	a := float32(1.0)
	b := 1.0
	{
		i := 0
		for true {
			a /= 10
			if a == 0 {
				fmt.Printf("%d,%s\n", i, fmt.Sprint(a))
				break
			}
			i++
		}
	}
	{
		i := 0
		for true {
			b /= 10
			if b == 0 {
				fmt.Printf("%d,%.10f", i, b)
				break
			}
			i++
		}
	}
}
