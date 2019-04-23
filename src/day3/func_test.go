package string_test

import (
	"math/rand"
	"testing"
	"fmt"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
	c, _ := returnMultiValues()
	t.Log(c)
}

func sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(sum(1,2,3,4,5))
}
func Clear() {
	fmt.Println("Clear resources ~")
}
func TestDefer(t *testing.T){
	defer Clear()
	fmt.Println("Start")
	panic("err")
}
