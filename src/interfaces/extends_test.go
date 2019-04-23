package interfaces_test

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

//type Dog struct {
//	p *Pet
//}
//
//func (d *Dog) Speak() {
//	fmt.Println("Dog~")
//}
//
//func (d *Dog) SpeakTo(host string) {
//	d.Speak()
//	fmt.Println(" ", host)
//}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Println("Dog~")
}


func TestDog(t *testing.T) {
	//1.go语言不支持显示的类型转换，
	//var p *Pet:=(*Pet)new(Dog)//这里编译器会编译出错
	dog := new(Dog)
	dog.SpeakTo("Tea")
}
