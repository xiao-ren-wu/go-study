package slice_test

import "testing"

func TestSliceInt(t *testing.T) {
	var s0 []int
	t.Log(len(s0),cap(s0))
	s0 = append(s0,1)

}

func TestSliceComparing(t *testing.T){
	a:=[...]int{1,2,3,4}
	b:=[...]int{1,2,3,4}
	if a==b {
		temp:=a[1:]
		t.Log(temp)
		t.Log("equals")
	}
}