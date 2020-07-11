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

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "testing")
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
func listAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "查询所有记录\n")
	list, _ := a.ListAll()
	for _, i := range list {
		fmt.Fprintf(w, "id为%d的name为%s,年龄为%s\n", i.Id, i.Name, i.Age)
	}
}
func selectById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "根据id查找记录\n")
	err := r.ParseForm()
	//res, _:=ioutil.ReadAll(r.Body) // 解析参数，默认是不会解析的
	//fmt.Println("______________",res)
	//fmt.Println(string(res))
	if err != nil {
		log.Fatal("parse form error ", err)
	}
	// 初始化请求变量结构 里面包含了json存储数据的各种可能
	formData := make(map[string]interface{})
	// 调用json包的解析，解析请求body
	json.NewDecoder(r.Body).Decode(&formData)
	for key, value := range formData {
		log.Println("key:", key, " => value :", value)
	}
	// 返回json字符串给客户端
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(formData)
}
func getJson2(w http.ResponseWriter, r *http.Request) {
	println("____________________进入getJson2________________")
	jsonData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("parse form error ", err)
	}
	// 初始化请求变量结构
	var formData map[string]interface{} //interface{}
	//结果与map[string]interface()结果是一样的
	// 调用json包的解析，解析请求body
	json.Unmarshal(jsonData, &formData)
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
	json.NewDecoder(r.Body).Decode(&formData)
	for key, value := range formData {
		log.Println("key:", key, " => value :", value)
	}
}
func Test_web1(t *testing.T) {
	//设置要解析的URL路由
	//相当配置java controller中的request maping
	http.HandleFunc("/hello", responseHello)
	http.HandleFunc("/test", test)
	http.HandleFunc("/listAll", listAll)
	http.HandleFunc("/cmd", crubTest)
	http.HandleFunc("/select", selectById)
	http.HandleFunc("/testJson", getJson)
	http.HandleFunc("/testJson2", getJson2)
	//设置监听的端口，开始监听
	errInfo := http.ListenAndServe(":8080", nil)
	if errInfo != nil {
		log.Fatal("ListenAndServe: ", errInfo)
	}

}
