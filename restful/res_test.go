package restful

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

type Item struct {
	Seq    int
	Result map[string]int
}

type Message struct {
	Dept    string
	Subject string
	Time    int64
	Detail  []Item
}

func getJson() ([]byte, error) {
	pass := make(map[string]int)
	pass["x"] = 50
	pass["c"] = 60
	item1 := Item{100, pass}

	reject := make(map[string]int)
	reject["l"] = 11
	reject["d"] = 20
	item2 := Item{200, reject}

	detail := []Item{item1, item2}
	m := Message{"IT", "KPI", time.Now().Unix(), detail}
	return json.MarshalIndent(m, "", "")

}

func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := getJson()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(resp))
}

func Test_res(t *testing.T) {
	//充当服务端,来进行json的生成与返回
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8085", nil)
}
