/*@Time : 2020/12/26 2:45 下午
@Author : ccc
@File : Interface1_test
@Software: GoLand*/
package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/lifei6671/gocaptcha"
	_ "github.com/vhaoran/vchat/common/g"
	"github.com/vhaoran/vchat/lib"
	"github.com/vhaoran/vchat/lib/ylog"
)

const debug = true

//const debug = false
//只有当接口值的
func init() {
	//------------ prepare modules----------
	//本步骤主要是装入系统必备的模块
	_, _ = lib.InitModulesOfOptions(&lib.LoadOption{
		LoadMicroService: false, //这不同必需要的
		LoadEtcd:         false, //etcd必須開啟，否則無法自動發現服務
		LoadPg:           false,
		LoadRedis:        false,
		LoadMongo:        false,
		LoadMq:           false,
		LoadJwt:          false,
		LoadES:           false,
	})
}
func TestInterface(t *testing.T) {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer) // enable collection of output  启用输出收集
	}
	f(buf) // NOTE: subtly incorrect!  巧妙的错误?
	if debug {
		// ...use buf...
		fmt.Println(buf)
	}
}

// If out is non-nil, output will be written to it.
//如果参数非空,输出将被写入它
func f(out io.Writer) {
	// ...do something...
	if out != nil {
		_, _ = out.Write([]byte("done!\n")) //if out is nil, here will panic
	}
}

func TestErr(t *testing.T) {
	e := errors.New("脑子进煎鱼了")
	w := fmt.Errorf("快抓住：%w", e)
	w2 := fmt.Errorf("跑了,%w", w)
	fmt.Println(w)
	//errors.unwrap只能褪去一层包装
	fmt.Println(errors.Unwrap(w))
	fmt.Println(w2)
	fmt.Println(errors.Unwrap(errors.Unwrap(w2)))
	fmt.Println(errors.Is(w, e))
	fmt.Println(errors.Is(e, w))
}

func TestErr2(t *testing.T) {
	e := errors.New("脑子进煎鱼了")
	fmt.Println(e)
	w := fmt.Errorf("抓点紧:%w", e)
	fmt.Println(errors.Unwrap(w))
	//拆包装至最后一层时再拆及时空对象
	fmt.Println(errors.Unwrap(errors.Unwrap(w)))
	fmt.Println(errors.Unwrap(errors.Unwrap(errors.Unwrap(w))))
}
func TestRandomString(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		n := 10
		randBytes := make([]byte, n/2)
		rand.Read(randBytes)
		fmt.Printf("%x\n", randBytes)
	}
}

func TestRedis(t *testing.T) {
	re := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
	//re := redis.NewClusterClient(&redis.ClusterOptions{Addrs: []string{"127.0.0.1:6379"},
	//	MaxRedirects:   2,
	//	ReadOnly:       false,
	//	RouteByLatency: false,
	//	RouteRandomly:  false,
	//	ClusterSlots:   nil,
	//	OnNewNode:      nil})
	//_, err := yredis.X.Set("pzy", "panzhenying", time.Second*10).Result()
	_, err := re.Set("pzy", "panzhenying", 3*time.Second).Result()
	if err != nil {
		ylog.Debug("这里有错误", err)
	}
	v := []string{"zs", "ls", "ww"}
	_, err = re.Set("name", v[0], 10*time.Second).Result()
	if err != nil {
		ylog.Debug("获取name有误", err)
	}
	m := make(map[string]interface{}, 0)
	m["zs"] = "beijing"
	m["ls"] = "shanghai"
	m["ww"] = "guangzhou"
	_, err = re.HMSet("address", m).Result()
	if err != nil {
		ylog.Debug("获取map有误", err)
	}
	//向一个对象追加内容  如果是字符串,增加字符串长度及内容
	_, err = re.Append("name", v[1]).Result()
	_, err = re.Append("name", v[2]).Result()
	if err != nil {
		ylog.Debug("append err", err)
	}
	re.Get("name")
	re.Set("num", 1, 20*time.Second).Result()
	re.Append("num", fmt.Sprint(23)).Result()
}

func TestYanZhengMa(t *testing.T) {
	//初始化一个验证码对象
	captchaImage := gocaptcha.NewCaptchaImage(1, 2, gocaptcha.RandLightColor())

	//画上三条随机直线
	//captchaImage.DrawLine(1)

	//画边框
	captchaImage.DrawBorder(gocaptcha.ColorToRGB(0x17A7A7A))

	//画随机噪点
	captchaImage.DrawNoise(gocaptcha.CaptchaComplexHigh)

	//画随机文字噪点
	captchaImage.DrawTextNoise(gocaptcha.CaptchaComplexLower)
	//画验证码文字，可以预先保持到Session种或其他储存容器种
	captchaImage.DrawText(gocaptcha.RandText(4))

	//将验证码保持到输出流种，可以是文件或HTTP流等
	//captchaImage.SaveImage(w, gocaptcha.ImageFormatJpeg)
	fmt.Println(gocaptcha.ImageFormatJpeg)
}
func TestRandom(t *testing.T) {
	num := 3
	txtChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	text := ""
	r := rand.New(rand.NewSource(time.Now().Unix()))
	textNum := len(txtChars)
	for i := 0; i < num; i++ {
		text = text + string(txtChars[r.Intn(textNum)])
	}
	fmt.Println(text)
}

func TestRandom2(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Printf("%.2f", float64(r.Intn(7))/float64(9))
}

func TestValueOf(t *testing.T) {
	l := []string{"zs", "ls", "ww"}
	//返回参数的值,如果参数为空,则返回0值
	fmt.Println("结果是:", reflect.ValueOf(l))
	//如果传入参数为指针,返回其指向的值,如果是空指针,间接返回一个0值,如果参数不是指针,则间接返回v
	fmt.Println("间接是:", reflect.Indirect(reflect.ValueOf(l)))
}

func TestBool(t *testing.T) {
	fmt.Println("zs" < "ls")
}

//sort interface排序
type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func TestSort(t *testing.T) {
	//序列排序
	names := []string{"zs", "ls", "ww"}
	sort.Strings(names)
	ylog.DebugDump(names)
	sort.Sort(StringSlice(names))
	ylog.DebugDump(names)
}

func TestErrInit(t *testing.T) {
	err := errors.New("this is an error")
	err1 := fmt.Errorf("%v, 请尽快处理", err)
	ylog.Debug("拆掉一层后", errors.Unwrap(err1))
	ylog.Error(err)
	ylog.Error(err1)
	ylog.Debug("拆掉一层后", errors.Unwrap(err))
}

func TestAssertions(t *testing.T) {
	//断言
	//一个类型断言检查它操作对象的动态类型是否和断言的类型匹配。
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File) // success: f == os.Stdout
	fmt.Println(f.Name())
	//c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
	//fmt.Println(c)
}
