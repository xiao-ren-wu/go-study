package string_test

import "testing"

func TestString(t *testing.T){
	var s string

	t.Log(s)//初始化默认值为零值""

	s="hello"

	t.Log(len(s))

	//s[1]='3' //string是不可以变类型的byte slice
	s="\xE4\xB8\xA5" //可以存储任何二进制数据

	t.Log(s)
	s="中"
	t.Log(len(s))//是byte数
	//取出byte slice数组中的指定位置的元素
	c:=[]rune(s)

	t.Logf("中 unicode %x",c[0])
	t.Logf("中 UTF8 %x",s)

}

