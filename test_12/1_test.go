/*@Time : 2020/12/7 4:00 下午
@Author : ccc
@File : 1.test
@Software: GoLand*/
package main

import (
	"testing"

	"github.com/vhaoran/vchat/lib/ylog"
)

func TestYlog1(t *testing.T) {
	ylog.Debug("123")
	println("456")
	ylog.Info("123")
}
