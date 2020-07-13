package restful

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func Test_res2(t *testing.T) {
	//充当客户端,来进行获取发送来的json数据,并进行解析,次文件相当于用浏览器访问"http://localhost:8085"
	url := "http://localhost:8085"
	ret, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer ret.Body.Close()

	body, err := ioutil.ReadAll(ret.Body)
	if err != nil {
		panic(err)
	}

	var msg Message
	err = json.Unmarshal(body, &msg)
	if err != nil {
		panic(err)
	}

	strTime := time.Unix(msg.Time, 0).Format("2006-01-02 15:04:05")
	fmt.Println("Dept:", msg.Dept)
	fmt.Println("Subject:", msg.Subject)
	fmt.Println("Time:", strTime, "\n", msg.Detail)
}
