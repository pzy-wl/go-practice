package t01

import (
	"fmt"
	"testing"
)

// go test -run=Test_int
func Test_int(t *testing.T) {
	a := 0
	b := int64(1)
	c := int8(1)
	d := int16(1)
	e := int32(40)
	fmt.Println("---", b, "---", a, c, d, e)
}
