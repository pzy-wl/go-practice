/*@Time : 2020/12/31 9:21 上午
@Author : ccc
@File : main
@Software: GoLand*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

//
//func Get(w http.ResponseWriter, r *http.Request) {
//	//初始化一个验证码对象
//	dx := 150
//	dy := 50
//	captchaImage := gocaptcha.NewCaptchaImage(dx, dy, gocaptcha.RandLightColor())
//
//	//画上三条随机直线
//	captchaImage.DrawLine(3)
//
//	//画边框
//	captchaImage.DrawBorder(gocaptcha.ColorToRGB(0x17A7A7A))
//
//	//画随机噪点
//	captchaImage.DrawNoise(gocaptcha.CaptchaComplexHigh)
//
//	//画随机文字噪点
//	captchaImage.DrawTextNoise(gocaptcha.CaptchaComplexLower)
//	//画验证码文字，可以预先保持到Session种或其他储存容器种
//	text := gocaptcha.RandText(4)
//	captchaImage.DrawText(text)
//	//将验证码保持到输出流种，可以是文件或HTTP流等
//	captchaImage.SaveImage(w, gocaptcha.ImageFormatJpeg)
//}
//func main() {
//	//生成验证码
//	err := gocaptcha.ReadFonts("fonts", ".ttf")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	http.HandleFunc("/get", Get) // 设置访问的路由
//
//	err = http.ListenAndServe(":9020", nil) // 设置监听的端口
//	if err != nil {
//		log.Fatal("ListenAndServe: ", err)
//	}
//}

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//for item, price := range db {
	//	fmt.Fprintf(w, "%s: %s\n", item, price)
	//}
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}
