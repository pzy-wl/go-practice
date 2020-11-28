/*@Time : 2020/10/21 6:53 下午
@Author : ccc
@File : main
@Software: GoLand*/
package main

import (
	"fmt"
)

func main() {
	//沸点和结冰点
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
