/**
  @author:pzy
  @date:2020/10/19
  @note:
**/
package cake

import (
	"testing"
	"time"
)

//此文件使用的测试方法为基准测试
//基准函数会运行目标代码 b.N 次。在基准执行期间，会调整 b.N 直到基准测试函数运行时间稳定
var defaults = Shop{
	//Verbose:      testing.Verbose(),
	Cakes:        20,
	BakeTime:     10 * time.Millisecond,
	NumIcers:     1,
	IceTime:      10 * time.Millisecond,
	InscribeTime: 10 * time.Millisecond,
}

func Benchmark(b *testing.B) {
	// Baseline: one baker, one icer, one inscriber.
	//一个烘焙师 一个糖霜师 一个镶嵌是
	// Each step takes exactly 10ms.  No buffers.
	//每一个步骤用时10ms,没有缓存区
	cakeshop := defaults
	cakeshop.Work(b.N) // 224 ms
}

func BenchmarkBuffers(b *testing.B) {
	// Adding buffers has no effect.
	// 增加缓存但是没有效果
	cakeshop := defaults
	cakeshop.BakeBuf = 10
	cakeshop.IceBuf = 10
	cakeshop.Work(b.N) // 224 ms
}

func BenchmarkVariable(b *testing.B) {
	// Adding variability to rate of each step
	//增加变量区改变每一步的时间
	// increases total time due to channel delays.
	//增加总时间,因为管道延时
	cakeshop := defaults
	cakeshop.BakeStdDev = cakeshop.BakeTime / 4
	cakeshop.IceStdDev = cakeshop.IceTime / 4
	cakeshop.InscribeStdDev = cakeshop.InscribeTime / 4
	cakeshop.Work(b.N) // 259 ms
}

func BenchmarkVariableBuffers(b *testing.B) {
	// Adding channel buffers reduces
	//使用的管道缓存变少了
	// delays resulting from variability.
	//延时因为变异性
	cakeshop := defaults
	cakeshop.BakeStdDev = cakeshop.BakeTime / 4
	cakeshop.IceStdDev = cakeshop.IceTime / 4
	cakeshop.InscribeStdDev = cakeshop.InscribeTime / 4
	cakeshop.BakeBuf = 10
	cakeshop.IceBuf = 10
	//cakeshop.Work(b.N) // 244 ms
	cakeshop.Work(4) // 244 ms
}

func BenchmarkSlowIcing(b *testing.B) {
	// Making the middle stage slower
	//使中间阶段变慢(霜糖时间增加)
	// adds directly to the critical path.
	//	直接添加关键路径
	cakeshop := defaults
	cakeshop.IceTime = 50 * time.Millisecond
	cakeshop.Work(b.N) // 1.032 s
}

func BenchmarkSlowIcingManyIcers(b *testing.B) {
	// Adding more icing cooks reduces the cost of icing
	//增减更多的初始用于减少霜糖使用的时间
	// to its sequential component, following Amdahl's Law.
	//根据Amdahl定律,对他的顺序进行区分----阿姆达尔定律:系统中某一部件
	//由于采用某种更快的执行方式后整个系统性能的提高与这种执行方式的使用频率或占总执行时间的比例有关
	cakeshop := defaults
	cakeshop.IceTime = 50 * time.Millisecond
	cakeshop.NumIcers = 5
	cakeshop.Work(b.N) // 288ms
}
