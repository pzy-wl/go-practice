package Web_server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net"
	"net/http"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func show(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func Test_web2(t *testing.T) {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	router.Run(":8000")
}
func Test_web3(t *testing.T) {
	//启动一个RESTFul服务
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello, ", html.EscapeString(request.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))

}

//定义自己的路由器
type MyMux struct {
}
type point struct {
	Metric    string            `json:"metric"`
	TimeStamp int64             `json:"timestamp"`
	Value     float64           `json:"value"`
	Tags      map[string]string `json:"tags"`
}

type responseExample struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func (mux *MyMux) sayHello(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "the method is not allowed！", http.StatusMethodNotAllowed)
		return
	}
	_, err := fmt.Fprintf(w, "Hello World!")
	fmt.Printf("在控制台输出Hello,World!\n")
	if err != nil {
		http.Error(w, "控制台打印出错1", http.StatusExpectationFailed)
		return
	}
}

func (mux *MyMux) sayHi(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "the method is not allowed！", http.StatusMethodNotAllowed)
		return
	}
	_, err := fmt.Fprintf(w, "Hi World!")
	fmt.Printf("在控制台输出Hi,World!\n")
	if err != nil {
		http.Error(w, "控制台打印出错2", http.StatusExpectationFailed)
		return
	}
}

func (mux *MyMux) writeJsonToClient(w http.ResponseWriter, r *http.Request) {

	response := make([]*responseExample, 3)
	//初始化response的信息
	for index, _ := range response {
		response[index] = &responseExample{
			Message: "this is test message " + strconv.Itoa(index),
			Error:   "this is test error " + strconv.Itoa(index),
		}
	}
	fmt.Println(r.URL)
	//解析go结构体变量为json数据
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("JSON marshaling  failed: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	//打印到控制台
	fmt.Printf("%s\n", jsonData)
	//打印到浏览器
	_, _ = fmt.Fprintf(w, "%s\n", jsonData)
}

func (muc *MyMux) parseJsonFromClient(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "the method is not allowed", http.StatusMethodNotAllowed)
	}

	if r.Header.Get("Content-TyPe") != "application/json" {
		http.Error(w, "please set Encode method application/json", http.StatusBadRequest)
	}
	//流式解码器
	br := bufio.NewReader(r.Body)

	// 查看前1个字节
	f, err := br.Peek(1)
	if err != nil || len(f) != 1 {
		http.Error(w, "peek error: "+err.Error(), http.StatusBadRequest)
		return
	}
	// Peek to see if this is a JSON array.
	var multi bool
	switch f[0] {
	case '{':
		multi = false
	case '[':
		multi = true
	default:
		http.Error(w, "expected JSON array or hash", http.StatusBadRequest)
		return
	}

	dps := make([]point, 1)
	//如果是多个对象{}组成的数组[],那么久解码到dps
	if dec := json.NewDecoder(br); multi {
		if err := dec.Decode(&dps); err != nil {
			http.Error(w, "json array decode error", http.StatusBadRequest)
			return
		}
	} else { //否则，解码到dps[0]
		if err = dec.Decode(&dps[0]); err != nil {
			http.Error(w, "json object decode error", http.StatusBadRequest)
			return
		}
	}
	// fmt.Print(len(dps))
	// fmt.Print(dps)
	jsonData, err := json.Marshal(dps)
	if err != nil {
		log.Fatalf("JSON marshaling  failed: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("%s\n", jsonData)
	_, _ = fmt.Fprintf(w, "%s\n", jsonData)
	//_, _ = fmt.Fprint(w, jsonData)
}

//实现http.Handler这个接口的唯一方法
func (mux *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path
	switch urlPath {
	case "/hello":
		mux.sayHello(w, r)
	case "/hi":
		mux.sayHi(w, r)
	case "/jsonPost":
		mux.parseJsonFromClient(w, r)
	case "/jsonGet":
		mux.writeJsonToClient(w, r)
	default:
		http.Error(w, "没有此url路径", http.StatusBadRequest)
	}
}

func Test_web4(t *testing.T) {
	//实例化路由器Handler
	mymux := &MyMux{}
	//基于TCP服务监听8088端口
	ln, err := net.Listen("tcp", ":8088")
	if err != nil {
		fmt.Printf("设置监听端口出错...")
	}
	//调用http.Serve(l net.Listener, handler Handler)方法，启动监听
	err1 := http.Serve(ln, mymux)
	if err1 != nil {
		fmt.Printf("启动监听出错")
	}
}
