/*@Time : 2020/12/9 5:28 下午
@Author : ccc
@File : cal_test
@Software: GoLand*/
package main

import (
	"testing"

	"github.com/vhaoran/vchat/lib/ylog"
)

//算法

//求解字符串出现的位置
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	var i, j int
	// i不需要到len-1
	for i = 0; i < len(haystack)-len(needle)+1; i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				//直接省去循环内的比较,进入下一次循环
				break
			}
		}
		// 判断字符串长度是否相等
		if len(needle) == j {
			return i
		}
	}
	return -1
}
func TestCal1(t *testing.T) {
	//给定两个字符串,求出第二个字符串在第一个字符串中出现的位置,如果没有则取值为-1
	println(strStr("hello, world", "wol"))
}

//回溯法
func subsets(nums []int) [][]int {
	// 保存最终结果
	result := make([][]int, 0)
	// 保存中间结果
	list := make([]int, 0)
	backtrack(nums, 0, list, &result)
	return result
}

// nums 给定的集合
// pos 下次添加到集合中的元素位置索引
// list 临时结果集合(每次需要复制保存)
// result 最终结果
func backtrack(nums []int, pos int, list []int, result *[][]int) {
	// 把临时结果复制出来保存到最终结果
	ans := make([]int, len(list))
	copy(ans, list)
	*result = append(*result, ans)
	// 选择、处理结果、再撤销选择
	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		backtrack(nums, i+1, list, result)
		list = list[0 : len(list)-1]
	}
}

func TestCal2(t *testing.T) {
	//使用回溯法求某个整数数组的所有可能子集(幂集)
	//(类似于0-1背包问题)---类似于完全二叉树简化为0-1问题
	l := []int{1, 2, 3, 4, 5, 6}
	ylog.DebugDump("数量是:", len(subsets(l)), subsets(l))
}
