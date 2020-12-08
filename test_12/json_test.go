/*@Time : 2020/12/8 9:36 上午
@Author : ccc
@File : json_test
@Software: GoLand*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

//测试json
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}

func TestJson1(t *testing.T) {
	//编码结构体数组为json,并将其输出
	data, err := json.Marshal(movies)
	if err != nil {
		return
	}
	fmt.Printf("%s\n", data)

}
func TestJson2(t *testing.T) {
	//有格式编码并输出, 两个参数分别是每行的前缀和每一层级的缩进
	data, err := json.MarshalIndent(movies, " ", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}
func TestJson3(t *testing.T) {
	//json解码 只对某些结构体成员进行解码 此例只填充title的值,其他json成员被忽略
	data, err := json.MarshalIndent(movies, " ", "  ")
	if err != nil {
		println("编码失败")
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		println("解码失败")
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
}

//*********************************************************************************8
const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func TestJson5(t *testing.T) {

}
