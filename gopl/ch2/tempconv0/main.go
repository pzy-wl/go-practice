/*@Time : 2020/10/22 9:46 上午
@Author : ccc
@File : main
@Software: GoLand*/
package main

import (
	"fmt"
)

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	//绝对0度
	AbsoluteZeroC Celsius = -273.15
	//结冰点
	FreezingC Celsius = 0
	//沸点
	BoilingC Celsius = 100
)

//摄氏度转华氏度
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

//华氏度转摄氏度
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
