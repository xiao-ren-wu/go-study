package object_pool

import (
	"time"
	"errors"
	"testing"
	"fmt"
)

type ReusableObj struct {
}

type ObjPool struct {
	bufChan chan *ReusableObj
}

func NewObjPool(numberOfObj int) *ObjPool {
	ObjPool := ObjPool{}
	ObjPool.bufChan = make(chan *ReusableObj, numberOfObj)
	for i := 0; i < numberOfObj; i++ {
		ObjPool.bufChan <- &ReusableObj{}
	}
	return &ObjPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}
}

func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}


func TestObjPool(t *testing.T){
	pool:=NewObjPool(10)
	for i:=0;i<11;i++{
		if v,err:=pool.GetObj(time.Second*1); err!=nil{
			t.Error(err)
		}else{
			fmt.Printf("%T\n",v)
		}
	}
}





















