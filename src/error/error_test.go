package main

import (
	"errors"
	"testing"
	"fmt"
)

//func getError() error {
//	return errors.New("I am a error ")
//}
//
//
//func TestGetError(t *testing.T) {
//	t.Log(getError())
//}
//
//
//func TestPanicVsExit(t *testing.T){
//
//	defer func() {
//		fmt.Println("Finally")
//	}()
//
//	fmt.Println("Start")
//
//	//panic(errors.New("Something wrong !"))
//
//	os.Exit(-1)
//
//}

func TestRecover(t *testing.T){
	defer func() {
		if err := recover(); err!=nil{
			fmt.Println("recover from ",err)
		}
	}()

	fmt.Println("Start")
	panic(errors.New("Something wrong !"))
}




























