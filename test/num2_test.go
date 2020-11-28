package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"testing"
)

func TestNum(t *testing.T) {
	lee1()
}
func num2(I float32) float32 {
	//阶梯工资
	var bonus float32 = 0.0
	switch {
	case I > 1000000:
		bonus = (I - 1000000) * 0.01
		I = 1000000
		fallthrough
	case I > 600000:
		bonus += (I - 600000) * 0.015
		I = 600000
		fallthrough
	case I > 400000:
		bonus += (I - 400000) * 0.03
		I = 400000
		fallthrough
	case I > 200000:
		bonus += (I - 200000) * 0.05
		I = 200000
		fallthrough
	case I > 100000:
		bonus += (I - 100000) * 0.075
		I = 100000
		fallthrough
	default:
		bonus += I * 0.1
	}
	fmt.Printf("提成总计：%f\n", bonus)
	return bonus
}
func num3(y, m, d int) {
	//：输入某年某月某日，判断这一天是这一年的第几天？
	var days int = 0
	switch m {
	case 12:
		days += d
		d = 30
		fallthrough
	case 11:
		days += d
		d = 31
		fallthrough
	case 10:
		days += d
		d = 30
		fallthrough
	case 9:
		days += d
		d = 31
		fallthrough
	case 8:
		days += d
		d = 31
		fallthrough
	case 7:
		days += d
		d = 30
		fallthrough
	case 6:
		days += d
		d = 31
		fallthrough
	case 5:
		days += d
		d = 30
		fallthrough
	case 4:
		days += d
		d = 31
		fallthrough
	case 3:
		days += d
		d = 28
		if (y%400 == 0) || (y%4 == 0 && y%100 != 0) {
			d += 1
		}
		fallthrough
	case 2:
		days += d
		d = 31
		fallthrough
	case 1:
		days += d
	}
	fmt.Printf("是今年的第 %d 天!\n", days)
}
func num5_sort() {
	//三个数字排序
	i := make([]int, 3)
	fmt.Scanf("%d%d%d", &i[0], &i[1], &i[2])
	sort.Ints(i)
	for j, v := range i {
		println(j, v)
	}
}
func num6_printC() {
	//打印出字母c使用*
}
func num8() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		println()
	}
}
func num11() {
	//	兔子总数问题  斐波那契数列
	var m1, m2 int = 1, 1

	for i := 1; i < 12; i++ {
		fmt.Println(m1, m2)
		m1 += m2
		m2 += m1
	}
}
func num12() {
	//	求素数  用一个大于等于2且小于等于这个数的开方的数去除这个数,如果能整除则说明不是素数,否则为素数
	var i, j, k, count int = 0, 0, 0, 0
	for i = 101; i <= 200; i++ {
		k = int(math.Sqrt(float64(i)))
		for j = 2; j <= k; j++ {
			if i%j == 0 {
				break
			}
		}
		if j == k+1 {
			fmt.Print(i, " ")
			count++
		}
	}
	fmt.Print("\n")
	fmt.Println("total", count)
}
func num13() {
	//	水仙花数 一个三位数各个位上的数字的立方想加等于这个数,则此数字为水仙花数
	for num := 100; num < 1000; num++ {
		i := num / 100
		j := num / 10 % 10
		k := num % 10
		if i*i*i+j*j*j+k*k*k == num {
			fmt.Printf("%d^3 + %d^3 + %d^3 = %d\n", i, j, k, num)
		}
	}
}
func num14() {
	//分解质因数
	var n, i int = 0, 0

	fmt.Printf("please input a number:")
	fmt.Scanf("%d\n", &n)
	fmt.Printf("%d = ", n)

	for i = 2; i <= n; i++ {
		for n != i {
			if n%i == 0 {
				fmt.Printf("%d * ", i)
				//层层减码
				n = n / i
			} else {
				break
			}
		}
	}

	fmt.Printf("%d\n", n)
}

type B bool

func num15() {
	//根据分数划分等级
	var score int = 0

	fmt.Printf("请输入分数：")
	fmt.Scanf("%d", &score)

	grade := B(score >= 90).C("A", B(score >= 60).C("B", "C"))
	fmt.Println("Grade =", grade)
}

//todo
func (b B) C(t, f interface{}) interface{} {

	//( a > b ) ? a : b 这是条件运算符的基本例子。 需要仔细参考
	if bool(b) == true {
		return t
	}
	return f
}
func num16() {
	//	公约数公倍数 辗转相处法
	var m, n, r, x int

	fmt.Print("请输入两个数：")

	fmt.Scanf("%d%d", &m, &n)
	x = m * n
	for n != 0 {
		r = m % n
		m, n = n, r
		//n = r
	}
	fmt.Printf("最大公约数：%d， 最小公倍数：%d\n", m, x/m)
}
func num17() {
	//	求字符串统计各类字符个数个数, 使用switch语句
	var i, j, k, l = 0, 0, 0, 0

	fmt.Print("请输入一串字符：")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	for _, ch := range input {
		switch {
		case ch >= 'A' && ch <= 'Z':
			i++
		case ch >= 'a' && ch <= 'z':
			i++
		case ch == ' ' || ch == '\t':
			j++
		case ch >= '0' && ch <= '9':
			k++
		default:
			l++
		}
	}
	fmt.Printf("char count = %d，space count = %d，digit count = %d，others count =%d\n", i, j, k, l)
}
func num19() {
	//	求出一千以内的所有的完数  一个数如果恰好等于它的因子之和，这个数就称为 “完数”。例如 6=1＋2＋3
	for i := 1; i <= 1000; i++ {
		num := 0
		for j := 1; j < i; j++ {
			if i%j == 0 {
				num += j
			}
		}
		if num == i {
			println(i)
		}
	}
}
func num20() {
	//	小球自由落体 一球从100 米高度自由落下，每次落地后反跳回原高度的一半；再落下，求它在第 10 次落地时，共经过多少米？第 10 次反弹多高？
	s := 100.0
	h := s / 2
	for i := 2; i <= 10; i++ {
		s += 2 * h
		h /= 2
	}
	fmt.Printf("总距离： %f\n", s)
	fmt.Printf("最后一次高度： %f\n", h)
}
func num21() {
	//	猴子吃桃问题 猴子第一天摘下若干个桃子，当即吃了一半，还不瘾，又多吃了一个。第二天早上又将剩下的桃子吃掉一半，又多吃了一个。以后每天早上都吃了前一天剩下的一半零一个。到第 10 天早上想再吃时，见只剩下一个桃子了。求第一天共摘了多少。
	num := 1
	for i := 10; i > 1; i-- {
		num = (num + 1) * 2
	}
	println("第一天有桃子", num, "个")
}
func num22() {
	//	比赛找对手 作死 其实脑子一转圈就能想出来
	var i, j, k int16 /*i 是 a 的对手，j 是 b 的对手，k 是 c 的对手*/
	for i = 'x'; i <= 'z'; i++ {
		for j = 'x'; j <= 'z'; j++ {
			if i != j {
				for k = 'x'; k <= 'z'; k++ {
					if i != k && j != k {
						if i != 'x' && k != 'x' && k != 'z' {
							fmt.Printf("比赛对手： a--%c b--%c c--%c\n", i, j, k)
						}
					}
				}
			}
		}
	}
}
func num23() {
	//有一分数序列：2/1，3/2，5/3，8/5，13/8，21/13…求出这个数列的前 20 项之和。
	number := 20
	a, b, s := 2.0, 1.0, 0.0
	for n := 1; n <= number; n++ {
		s = s + a/b
		a, b = a+b, a
	}
	fmt.Printf("sum is %9.6f\n", s)
}
func num25() {
	//	求 1+2!+3!+…+20! 的和
	sum := 1
	for i := 2; i <= 20; i++ {
		num := 1
		for j := 2; j <= i; j++ {
			num *= j
		}
		sum += num
		println(sum)
	}
	println("阶乘是:", sum)
}
func lee1() {
	var nums = []int{2, 7, 11, 15}
	target := 9
	length := len(nums)
	match := make(map[int]int, length)
	for i := 0; i < length; i++ {
		//println(nums[i])
		if index, ok := match[target-nums[i]]; ok {
			println(index, i)
		}
		match[nums[i]] = i
		println(match[nums[i]])
	}
}
