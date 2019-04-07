package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	t.Log(arr[1], arr[2])
}
func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	for idx, e := range arr3 {
		t.Log(idx, e)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4}
	arr3Sec := arr3[:3]
	arr3Dec := arr3[2:]
	t.Log(arr3Sec, arr3Dec)
}
