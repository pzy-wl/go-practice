/**
  @author:pzy
  @date:2020/10/19
  @note:
**/
package cake

import (
	"fmt"
	"math/rand"
	"time"
)

type Shop struct {
	Verbose        bool
	Cakes          int           // number of cakes to bake  需要烘焙的蛋糕数量
	BakeTime       time.Duration // time to bake one cake  蛋糕的烘焙时间
	BakeStdDev     time.Duration // standard deviation of baking time  烘焙标准误差时间
	BakeBuf        int           // buffer slots between baking and icing  烘焙和糖霜之间的缓冲槽
	NumIcers       int           // number of cooks doing icing	 做糖霜的厨子数
	IceTime        time.Duration // time to ice one cake	为一个蛋糕做糖霜的时间
	IceStdDev      time.Duration // standard deviation of icing time	糖霜标准误差时间
	IceBuf         int           // buffer slots between icing and inscribing  糖霜和嵌入之间的缓冲槽
	InscribeTime   time.Duration // time to inscribe one cake	嵌入一个蛋糕的时间
	InscribeStdDev time.Duration // standard deviation of inscribing time	嵌入蛋糕的标准时间
}

type cake int

func (s *Shop) baker(baked chan<- cake) {
	for i := 0; i < s.Cakes; i++ {
		c := cake(i)
		if s.Verbose {
			fmt.Println("baking", c)
		}
		//烘焙工作时间 也相当于是休眠时间
		work(s.BakeTime, s.BakeStdDev)
		baked <- c
	}
	close(baked)
}

func (s *Shop) icer(iced chan<- cake, baked <-chan cake) {
	for c := range baked {
		if s.Verbose {
			fmt.Println("icing", c)
		}
		//糖霜工作时间
		work(s.IceTime, s.IceStdDev)
		iced <- c
	}
}

func (s *Shop) inscriber(iced <-chan cake) {
	for i := 0; i < s.Cakes; i++ {
		c := <-iced
		if s.Verbose {
			fmt.Println("inscribing", c)
		}
		//嵌入工作时间
		work(s.InscribeTime, s.InscribeStdDev)
		if s.Verbose {
			fmt.Println("finished", c)
		}
	}
}

// Work runs the simulation 'runs' times.
//	工作运行模拟运行的次数
func (s *Shop) Work(runs int) {
	for run := 0; run < runs; run++ {
		baked := make(chan cake, s.BakeBuf)
		iced := make(chan cake, s.IceBuf)
		go s.baker(baked)
		for i := 0; i < s.NumIcers; i++ {
			go s.icer(iced, baked)
		}
		s.inscriber(iced)
	}
}

// work blocks the calling goroutine for a period of time
// that is normally distributed around d
// with a standard deviation of stddev.
func work(d, stddev time.Duration) {
	delay := d + time.Duration(rand.NormFloat64()*float64(stddev))
	time.Sleep(delay)
}
