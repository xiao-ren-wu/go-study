package type_test

import "testing"

func TestPoint(t *testing.T){
	a:=1
	aPtr:=&a

	t.Log(a,aPtr,*aPtr)
	t.Logf("%T %T",a,aPtr)
}
