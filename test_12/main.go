/*@Time : 2020/12/3 11:13 上午
@Author : ccc
@File : main
@Software: GoLand*/
package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch4/github"
)

func main() {
	fmt.Println("hello where!")
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
