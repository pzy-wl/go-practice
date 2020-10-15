/**
  @author:pzy
  @date:2020/10/6
  @note:
**/
package main

import (
	"strings"
	"testing"
)

//用于测试字符串处理    --符合预期
func TestSplit(t *testing.T) {
	name := "http://p.0755yicai.com//Users/ccc/work/yicms/dao/img/uploadimages*article*picture*0*147*147733_20200704154937649_0.jpg?e=1600308810&token=gEpp05gnISRQeLZ6d5GCnAryXSFDnMfl_G5iG5p5:wxO3rtgTSXhd7tVv--QU-o3Psv8="
	name = strings.Split(name, "?")[0]
	println(name) // 输出结果http://p.0755yicai.com//Users/ccc/work/yicms/dao/img/uploadimages*article*picture*0*147*147733_20200704154937649_0.jpg
}
