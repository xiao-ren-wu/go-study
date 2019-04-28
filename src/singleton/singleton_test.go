package singleton__test

import (
	"sync"
	"fmt"
	"testing"
	"unsafe"
)

type Singleton struct {

}

var singleInstance *Singleton

var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("create a Obj")
		singleInstance = new (Singleton)
	})
	return singleInstance
}

func CreateSingletonObj() *Singleton {
	return GetSingletonObj()
}

func TestCreateSingletonObj(t *testing.T){
	var wg sync.WaitGroup
	for i:=0;i<10;i++{
		wg.Add(1)
		go func() {
			obj:=CreateSingletonObj()
			fmt.Printf("%d\n",unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
