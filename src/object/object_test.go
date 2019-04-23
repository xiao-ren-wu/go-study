package object_test

import (
	"testing"
	"fmt"
)

type Emloyee struct {
	Id   string
	Name string
	Age  int
}

func (e *Emloyee) String() string {
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d",e.Id,e.Name,e.Age)
}

//func (e Emloyee) String() string {
//	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
//}
//func TestCreateEmployeeObj(t *testing.T) {
//	e := Emloyee{"0", "Bob", 20}
//	e1 := Emloyee{Name: "Mike", Age: 30}
//	e2 := new(Emloyee) //返回指针
//	e2.Id = "2"
//	e2.Name = "Rose"
//	e2.Age = 20
//	e.String()
//	t.Log(e)
//	t.Log(e1)
//	t.Log(e2.Name)
//}
func TestStructOperations(t *testing.T) {
	e:=Emloyee{Name:"Tea",Age:18,Id:"12"}
	t.Log(e.String())
}
