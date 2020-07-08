package Web_server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func Test_restful2(t *testing.T) {
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
	fmt.Println("======执行至此!==============")
	err = json.Unmarshal(body, &msg)
	if err != nil {
		fmt.Println("手动panic!!!!")
		//panic(err)
	}

	strTime := time.Unix(msg.Time, 0).Format("2006-01-02 15:04:05")
	fmt.Println("Dept:", msg.Dept)
	fmt.Println("Subject:", msg.Subject)
	fmt.Println("Time:", strTime, "\n", msg.Detail)
}

/*
//运行结果:
Dept: IT
Subject: KPI
Time: 2015-02-28 16:43:11
 [{100 map[c:60 x:50]} {200 map[d:20 l:11]}]
*/
