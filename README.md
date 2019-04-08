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

### 数组和切片

##### 声明数组

- 关键字 数组名称 数组 ------>var arr [3]int
- 也可以不指定数组的大小，通过初始化来指定数组的大小
  - `arr3:=[...]int{1,2,3,4}`

##### 数组的遍历

- 传统的for循环

  ~~~go
  for i:=0;i<len(arr3);i++{
  	t.Log(arr3[i])
  }
  ~~~

- 使用关键字range

  ~~~go
  for idx,e:=range arr3{
  	t.Log(idx,e)
  }
  ~~~

  变量idx为数组下标，e为数组中下标为idx对应的值，如果不使用idx，可以使用‘_’替代。

##### 截取指定区间的元素值   

`a[开始索引（包含）:结束索引（不包含）]`

如果取前n或者后n个元素，可以省略对应边界值

~~~go
arr3 := [...]int{1, 2, 3, 4}
arr3Sec := arr3[:3]//获取前3各元素
arr3Dec := arr3[2:]//获取数组下标从2往后所有的元素
arr3All := arr3[:]//获取所有元素
~~~

#### 切片

##### 声明切片

~~~go
var s0 []int //切片的声明不指定数组的长度	
s1:=[...]int{1,2,3,4}
~~~



##### 内部结构

![](H:\go-study\images\切片.png)

go语言中的每一个数组都有一个对应的“数组头”的数据结构，类似redis中的SDS。

“数组头”包含数组的元素个数以及数组的容量，当当前指针指向的数组的容量不能满足时，便会重新申请一块长度为原数组长度两倍的容量，并将原数组的内容拷贝到新数组中。

当我们使用切片得到的新数组系统并不会给我们分配一块新的内存，只是给了我们数组指定地址的引用。所以如果一个数组有多个切片，更新其中一个数组的内容，其他的数组也会受影响。     

【如图】

![](H:\go-study\images\切片共享.png)



#### 切片VS数组

- 切片通过指针引用数组的形式，使得切片的长度可以伸缩
- 数组之间可以进行比较，切片不能进行比较

### Map

#### 声明map

```go
m:=map[string]int{"one":1,"two":2,"three":3}
m1:=map[string]int{}
m2:=make(map[string]int,10)
```

在go语言中，获取通过key获取map中指定的value时，返回两个返回值，第一个为value,另一个为bool类型，如果key存在，返回true,如果key不存在返回false。

在go语言中，如果访问的key不存在，go语言返回0,所以我们无法通过返回值是否为nil确定key是否存在。但是，我们可以通过获取key时返回的另一个返回值确定key是否存在。

```go
func TestKey(t *testing.T) {
   m := map[int]int{}
   if v, ok := m[23]; ok {
      t.Log("Key is exist",v)
   }else {
      t.Log("Key is not exist")
   }
}
```

#### 遍历map

```go
func TestTravelMap(t *testing.T) {
   m1:=map[int]int{1:2,2:3,4:6,5:6}
   for k,v :=range m1{
      t.Log(k,v)
   }
}
```

map的value还可以存储方法

```go
func TestMapWithFunValue(t *testing.T) {
   m:=map[int]func(op int)int{}
   m[1]= func(op int) int {
      return op*op
   }
   t.Log(m[1](2))
}
```

### 字符串

1. string是数据类型，不是指针或者引用类型
2. string是只读的byte slice,len函数可以包含它所包含的byte数
3. string的byte数组可以存放任何数据

eg:

```go
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
```

### 函数

**Go语言函数与其他语言的差异**

1. 函数可以返回多个值
2. 所有的参数都是值传递：slice,map,channel会引起传引用的错觉
3. 函数可以作为变量的值
4. 函数可以作为参数和返回值

```go
//定义有两个返回值的函数
func returnMultiValues() (int, int) {
   return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
    //变量a,b分别用来接收返回值
   a, b := returnMultiValues()
   t.Log(a, b)
    //如果想忽略第二个返回值，使用'_'替代即可
   c,_:=returnMultiValues()
   t.Log(c)
}
```

#### 可变长参数

```go
func sum(ops ...int) int {
   ret := 0
   for _, op := range ops {
      ret += op
   }
   return ret
}

func TestVarParam(t *testing.T) {
   t.Log(sum(1,2,3,4,5))
}
```

#### 延迟执行函数

`defer`函数：在函数返回前执行，一般用于回收某些资源或者释放某些锁等。

```go
func Clear() {
   fmt.Println("Clear resources ~")
}
func TestDefer(t *testing.T){
   defer Clear()
   fmt.Println("Start")
}
//输出
//Start
//Clear resources ~
//--- FAIL: TestDefer (0.00s)
//panic: err [recovered]
//	panic: err
```

### 面向对象



















































