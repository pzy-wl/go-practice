package t02

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func Test_err_1(t *testing.T) {
	//
	s1 := "0.123456"
	i, err := strconv.ParseFloat(s1, 64)
	fmt.Println(err)
	if err != nil {
		fmt.Println("-----------------", i)
		fmt.Println("-----------------", err)
	}
	s := err.Error()
	fmt.Println("-----------------", s)
}

func Test_panic(t *testing.T) {
	panic("this is an error")
}

func Test_defer(t *testing.T) {
	err := f(3)
	if err != nil {
		fmt.Println("err_test->", err)
	}

	fmt.Println("-----good this pos------------")
}

func f(i int) (er1 error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("#########", err)
			er1 = errors.New(fmt.Sprint(err))
		}
	}()

	defer func() {
		fmt.Println("---defer-1--")
	}()

	defer func() {
		fmt.Println("---defer--2-")
	}()

	defer func() {
		fmt.Println("---defer-3--")
	}()

	fmt.Println("-----hello------------")
	panic(" ************* error ")
}
