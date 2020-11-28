/*@Time : 2020/10/21 6:51 下午
@Author : ccc
@File : main
@Software: GoLand*/
package main

import (
	"fmt"
)

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	// Output:
	// boiling point = 212°F or 100°C
}
