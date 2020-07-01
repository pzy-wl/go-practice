package t03

import (
	"fmt"
)

type (
	ABC struct {
		I int
	}
)

func Call() {
	//bean := new(ABC)
	good()
}

func good() {
	fmt.Println("aaa")
}
