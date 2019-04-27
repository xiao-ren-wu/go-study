package groutine__test

import (
	"testing"
	"sync"
)

//
//func TestGroutine(t *testing.T){
//	for i:=0;i<10;i++{
//		go func(i int) {
//			fmt.Println(i)
//		}(i)
//	}
//}

//func TestCounter(t *testing.T) {
//	counter := 0
//	for i := 0; i < 5000; i++ {
//		go func() {
//			counter++
//		}()
//	}
//	time.Sleep(1 * time.Second)
//	t.Logf("counter = %d", counter)
//}
func TestCounterThreadSafeWithWaitGroup(t *testing.T) {
	mut:=sync.Mutex{}
	wg:=sync.WaitGroup{}
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	//等待协程执行完，程序才退出
	wg.Wait()
	t.Logf("counter = %d", counter)
}