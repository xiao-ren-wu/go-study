package map_test

import "testing"

func TestMap(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log("len", len(m))
	m1 := map[string]int{}
	m1["id1"] = 3
	t.Log("len", len(m1))
	m2 := make(map[string]int, 10)
	t.Log("len m2", len(m2))
}
func TestKey(t *testing.T) {
	m := map[int]int{}
	if v, ok := m[23]; ok {
		t.Log("Key is exist",v)
	}else {
		t.Log("Key is not exist")
	}
}
func TestTravelMap(t *testing.T) {
	m1:=map[int]int{1:2,2:3,4:6,5:6}
	for k,v :=range m1{
		t.Log(k,v)
	}
}

func TestMapWithFunValue(t *testing.T) {
	m:=map[int]func(op int)int{}
	m[1]= func(op int) int {
		return op*op
	}
	t.Log(m[1](2))
}



















