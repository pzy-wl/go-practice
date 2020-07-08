package Web_server

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"
)

func test(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "testing")
}
func responseHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析参数，默认是不会解析的
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	//t0:=time.Now()
	name := GetRandomName()
	insert(name, "电商", time.Now())
	//fmt.Println("插入完成!",time.Since(t0))
	////selectAll()
	////将字符串发送给客户端
	fmt.Println("敲门!", time.Now())
	fmt.Fprintf(w, "Hello Golang!")
	//	客户端请求一次这个函数执行两次?
}
func Test_web1(t *testing.T) {
	//设置要解析的URL路由
	//相当配置java controller中的request maping
	http.HandleFunc("/hello", responseHello)
	http.HandleFunc("/test", test)
	//http.HandleFunc("www.baidu.com", responseHello)
	//设置监听的端口，开始监听
	errInfo := http.ListenAndServe(":8080", nil)
	if errInfo != nil {
		log.Fatal("ListenAndServe: ", errInfo)
	}

}
