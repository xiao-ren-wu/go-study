package object_pool_test

import (
	"testing"
	"sync"
	"fmt"
)

//func TestSyncPool(t *testing.T){
//	pool:=&sync.Pool{
//		New: func() interface{} {
//			fmt.Println("Create a new object.")
//			return 100
//		},
//	}
//	v:=pool.Get().(int)
//	fmt.Println(v)
//	pool.Put(3)
//	v1,_:=pool.Get().(int)
//	fmt.Println(v1)
//}
func TestSyncPoolMultiGroutine(t *testing.T){
	pool:=&sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new Obj")
			return 10
		},
	}
	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup

	for i:=0;i<10;i++{
		wg.Add(1)
		go func(id int) {
			t.Log(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
