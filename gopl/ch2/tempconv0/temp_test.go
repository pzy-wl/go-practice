/*@Time : 2020/10/22 9:48 上午
@Author : ccc
@File : temp_test
@Software: GoLand*/
package main

import (
	"fmt"
)

func Example_one() {
	//example是预设好输出结果,如果一直则为运行成功,否则是运行失败,不预设结果就会导致测试文件不能运行
	{
		//!+arith
		fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
		boilingF := CToF(BoilingC)
		fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
		//!-arith
	}

	// Output:
	// 100
	// 180
}

func Example_two() {

	c := FToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))

	// Output:
	// 100°C
	//100°C
	//100°C
	//100°C
	//100
	//100
}
