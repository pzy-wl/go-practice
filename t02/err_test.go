package t02

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func Test_err_1(t *testing.T) {
	//
	i, err := strconv.ParseFloat("aaaa", 64)

	//
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
	err := f()
	if err != nil {
		fmt.Println("err_test->", err)
	}

	fmt.Println("-----good this pos------------")
}

func f() (er1 error) {
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

	fmt.Println("-----hellow------------")
	panic(" ************* error ")
}
