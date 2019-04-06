# :name_badge: go-study

[TOC]



## hello World

~~~go
func main(){
	fmt.Print("hello world")
}
~~~

**应用程序入口[注意]**

- 必须是main包：package main
- 必须是main方法：func main()
- 文件名不一定是main.go

- Go中main函数不支持任何返回值，通过os.Exit来返回状态
- Go中main函数不支持入参，在程序中通过os.Args来获取命令行参数

### 测试程序编写

1. 源码文件以_test结尾：xxx_test.go
2. 测试方法名称以Test开头：func TestXXX(t  *testing.T){...}

eg:

~~~go
package test

import "testing"

func TestFirstTry(t *testing.T){
	t.Log("My first try!")
}
~~~

### 变量和常量

变量的定义：

~~~go
var a int = 1
var b int = 2
//go语言有变量类型自动推断能力，等价于
var a = 1
var b = 2
//等价于
var(
	a = 1
    b = 2
)
//等价于
a:=1
b:=2
~~~

go变量和其他编程语言的差异

- 复制可以自动类型推断
- 在一个赋值语句中可以对多个变量进行同时赋值

eg:

~~~go
//交换两个变量的值
func TestExchange(t *testing.T){
	a:=1
	b:=2
	a,b=b,a
}
~~~

常量

go语言的常量可以对指定步长或者指定移位运算进行简化

~~~go
const (
	Monday = iota+1
	Tuesday
	Wednesday
)
const(
	Readable=1<<iota
	Writable
	Executable
)
~~~

### go语言基本数据类型

| bool                                                   |
| ------------------------------------------------------ |
| string                                                 |
| int  int8 int16 int32 int64                            |
| uint uint8 uint16 uint32 uint64 uintptr                |
| byte //alias for uint8                                 |
| rune //alias for int32,repewsents a Unicode code point |
| float32 float64                                        |
| complex64 complex128                                   |

**[注意]**

- go语言不允许隐式类型转换
- 别名和原有类型也不能进行隐式类型转换

**指针**

go语言中不支持指针运算！！！

aPtr+=1    **×**

```go
func TestPoint(t *testing.T){
   a:=1
   aPtr:=&a
   t.Log(a,aPtr,*aPtr)
   t.Logf("%T %T",a,aPtr)
}
```

### 运算符

#### 算数运算符

| 运算符 | 描述 |       示例       |
| :----: | :--: | :--------------: |
|   +    | 相加 |  A+B输出结果30   |
|   -    | 相减 | A-B输出结果为-10 |
|   *    | 相乘 | A*B输出结果为200 |
|   /    | 相除 |  B/A输出结果为2  |
|   %    | 求余 |  B%A输出结果为0  |
|   ++   | 自增 | A++输出结果为11  |
|   --   | 自减 |  A--输出结果为9  |

**Go语言中没有前置的++，--**

#### 比较运算符

| 运算符 |                            描述                            |     实例      |
| :----: | :--------------------------------------------------------: | :-----------: |
|   ==   |     检查两个值是否相等，如果相等返回True否则返回False      | (A==B)为False |
|   !=   |   检查两个值是否不相等，如果不相等返回True否则返回False    | (A!=B)为True  |
|   >    |    检查左边是否大于右边值，如果是返回True,否则返回False    | (A>B)为False  |
|   <    |   检查左边值是否小于右边值，如果是返回True,否则返回False   |  (A<B)为True  |
|   >=   |  检查左边是否大于等于右边值，如果是返回True,否则返回False  | (A>=B)为False |
|   <=   | 检查左边值是否小于等于右边值，如果是返回True,否则返回False | (A<=B)为True  |

**用等号比较数组**

- 相同维数且含有相同个数的数组才可以比较
- 每个元素都相等才相等

```go
func TestCompareArray(t *testing.T){
   a:=[...]int{1,2,3,4}
   b:=[...]int{2,3,4,5}
   c:=[...]int{1,2,3,4,5}
   d:=[...]int{1,2,3,4}
   t.Log(a==b)
   t.Log(a==c)//Invalid operation: a==c (mismatched types [4]int and [5]int)
   t.Log(a==d)
}
//out:
false
true
```

#### 位运算符

&^按位置清零

~~~bash
1&^0--1
1&^1--0
0&^1--0
0&^0--0
~~~





















