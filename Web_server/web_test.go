package Web_server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/vhaoran/go-practice/t03"
)

var a = new(t03.AbcDao)

func test(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "testing")
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
	fmt.Println("敲门!", time.Now())
	fmt.Fprintf(w, "Hello Golang!")
}
func crubTest(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()   // 解析参数，默认是不会解析的
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Println("敲门!", time.Now())
	fmt.Fprintf(w, "Hello Golang!")
}
func listAll(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "查询所有记录\n")
	list, _ := a.ListAll()
	for _, i := range list {
		fmt.Fprintf(w, "id为%d的name为%s,年龄为%s\n", i.Id, i.Name, i.Age)
	}
}
func selectById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "根据id查找记录\n")
	err := r.ParseForm()
	if err != nil {
		log.Fatal("parse form error ", err)
	}
	// 初始化请求变量结构 里面包含了json存储数据的各种可能
	formData := make(map[string]int64)
	// 调用json包的解析，解析请求body
	_ = json.NewDecoder(r.Body).Decode(&formData)
	for key, value := range formData {
		log.Println("key:", key, " => value :", value)
		println("测试输出---循环内")
		ab, _ := a.Get(value)
		fmt.Println("查询到的数据是:", ab)
	}

	// 返回json字符串给客户端
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(formData)
}
func getJson2(_ http.ResponseWriter, r *http.Request) {
	println("____________________进入getJson2________________")
	jsonData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("parse form error ", err)
	}
	// 初始化请求变量结构
	var formData map[string]interface{} //interface{}
	//结果与map[string]interface()结果是一样的
	// 调用json包的解析，解析请求body
	_ = json.Unmarshal(jsonData, &formData)
	//fmt.Println(formData)
	if err != nil {
		fmt.Println(err)
		return
	}
	m := formData
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "type: string\nvalue: ", vv)
			fmt.Println("------------------")
		case float64:
			fmt.Println(k, "type: float64\nvalue: ", vv)
			fmt.Println("------------------")
		case bool:
			fmt.Println(k, "type: bool\nvalue: ", vv)
			fmt.Println("------------------")
		case map[string]interface{}:
			fmt.Println(k, "type: map[string]interface{}\nvalue: ", vv)
			for i, j := range vv {
				fmt.Println(i, ": ", j)
			}
			fmt.Println("------------------")
		case []interface{}:
			fmt.Println(k, "type: []interface{}\nvalue: ", vv)
			for key, value := range vv {
				fmt.Println(key, ": ", value)
				//如果是嵌套结构,继续解码并输出
				fmt.Println("检验是否有嵌套结构")
				//*.(type)只能在switch语句中使用
				switch vt := value.(type) {
				case map[string]int:
					for m, n := range vt {
						fmt.Println(m, ": 123", n)
					}
				default:
					fmt.Println(k, "type: nil\nvalue: ", vt)
					fmt.Println("------------------")
				}
			}
			fmt.Println("------------------")
		default:
			fmt.Println(k, "type: nil\nvalue: ", vv)
			fmt.Println("------------------")
		}
	}
}

func getJson(w http.ResponseWriter, r *http.Request) {
	println("____________________进入getJson________________")
	err := r.ParseForm()
	if err != nil {
		log.Fatal("parse form error ", err)
	}
	// 初始化请求变量结构
	formData := make(map[string]interface{})
	// 调用json包的解析，解析请求body
	_ = json.NewDecoder(r.Body).Decode(&formData)
	for key, value := range formData {
		log.Println("key:", key, " => value :", value)
	}
}

type myHandler struct{}

func (_ *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	w.Write([]byte("hello, this is myHandle!, requestUrl is " + r.URL.String()))
}

func Test_web1(t *testing.T) {

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:         ":8080",
		WriteTimeout: 4 * time.Second,
	}
	server.Handler = mux

	//跟路由包含所有的未注册的路由
	mux.Handle("/", &myHandler{})
	//设置要解析的URL路由
	//相当配置java controller中的request maping
	mux.HandleFunc("/hello", responseHello)
	mux.HandleFunc("/test", test)
	mux.HandleFunc("/listAll", listAll)
	mux.HandleFunc("/cmd", crubTest)
	mux.HandleFunc("/select", selectById)
	mux.HandleFunc("/testJson", getJson)
	mux.HandleFunc("/testJson2", getJson2)
	mux.HandleFunc("/toShow", show)
	//设置监听的端口，开始监听
	errInfo := server.ListenAndServe()
	//与上面的功能相同
	//errInfo := http.ListenAndServe(":8080", mux)
	if errInfo != nil {
		if errInfo == http.ErrServerClosed {
			log.Println("Server closed under request!")
		} else {
			log.Fatal("server closed unexpect!")
		}
	}
	log.Println("server exit!")

}
