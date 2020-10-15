/**
  @author:pzy
  @date:2020/10/12
  @note:
**/
package main

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	s := "不过，坦白说有时把爱情看成一项\"计划\"而实行，很难让你品尝到人世的真情，唯有真情才是爱情中最珍贵的。"
	println(s)
	println(strings.Contains(s, "\""))
	println(strings.Replace(s, "\"", "'", -1))
}
