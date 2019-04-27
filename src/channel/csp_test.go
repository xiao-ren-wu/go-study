package csp_test

import (
	"time"
	"fmt"
)

func service() string {
	time.Sleep(time.Millisecond*500)
	return "Done"
}

func AsyncService() chan string{
	retCh:=make(chan string,1)

	go func() {
		ret:=service()
		fmt.Println("returned result.")
		retCh<-ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func otherTask(){
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond*1000)
	fmt.Println("Task is done.")
}

//func TestService(t *testing.T)  {
//	fmt.Println(service())
//	otherTask()
//}
//
//func TestAsyncService(t *testing.T){
//	reCh:=AsyncService()
//	otherTask()
//	fmt.Println(<-reCh)
//}
