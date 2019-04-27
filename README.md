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

![](https://github.com/xiao-ren-wu/go-study/blob/master/images/%E5%88%87%E7%89%87.png)

go语言中的每一个数组都有一个对应的“数组头”的数据结构，类似redis中的SDS。

“数组头”包含数组的元素个数以及数组的容量，当当前指针指向的数组的容量不能满足时，便会重新申请一块长度为原数组长度两倍的容量，并将原数组的内容拷贝到新数组中。

当我们使用切片得到的新数组系统并不会给我们分配一块新的内存，只是给了我们数组指定地址的引用。所以如果一个数组有多个切片，更新其中一个数组的内容，其他的数组也会受影响。     

【如图】

![](https://github.com/xiao-ren-wu/go-study/blob/master/images/%E5%88%87%E7%89%87%E5%85%B1%E4%BA%AB.png)



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

### :man:面向对象

####  属性定义

go语言采用结构体来封装’属性‘

~~~go
type Emloyee struct {
	Id   string
	Name string
	Age  int
}
~~~

#### 行为定义

~~~go
//一般使用该种定义方式，采用指针的形式，不会复制变量，节约内存空间
func (e *Emloyee) String() string {
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d",e.Id,e.Name,e.Age)
}
//该种方式会额外复制一次变量的值，造成内存浪费。不推荐使用
func (e Emloyee) String() string {
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}
~~~

#### 实例化’对象‘

即实例化结构体

~~~go
func TestCreateEmployeeObj(t *testing.T) {
    //1.按照结构体丁定义属性的顺序依次赋值
	e := Emloyee{"0", "Bob", 20}
    //通过制定属性的形式为结构体赋值
	e1 := Emloyee{Name: "Mike", Age: 30}
    //返回指针类型
	e2 := new(Emloyee) 
	e2.Id = "2"
	e2.Name = "Rose"
	e2.Age = 20
	e.String()
	t.Log(e)
	t.Log(e1)
	t.Log(e2.Name)
}
~~~

#### 接口

1. 接口定义

~~~go
type Programmer interface{
    WriteHelloWorld() string
}
~~~

2. 接口实现

~~~go
type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"hello World\")"
}
~~~

3. test

~~~go
func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
//out:
//fmt.Println("hello World")
~~~

go语言并没有通过向java的关键字`implement`实现改接口，只是实现的函数和接口中的函数签名保持一致，就认为该接口中的方法被实现了，称为`Duck Type`

#### 接口变量

#### 自定义类型

自定义类型感觉像C语言中的`#define`一样，可以简化程序的书写。

eg:

~~~go
//将函数的签名进行替换
type IntConv func(op int) int
//通过使用自定义类型，简化函数的书写，在一定程度上也增加了程序的可读性
func timeSpent(inner IntConv)IntConv{
	return func(n int) int {
		start:=time.Now()
		ret:=inner(n)
		fmt.Println("time spent:",time.Since(start).Seconds())
		return ret
	}
}
~~~

#### ”多态“

1. 先定一个父类：

~~~go
type Pet struct {
}
func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}
~~~

2. 再定义一个子类：

   该子类’假装继承‘了Pet父类，并重写了父类的Speak方法

~~~go
type Dog struct {
	Pet
}
func (d *Dog) Speak() {
	fmt.Println("Dog~")
}
~~~

3. test

~~~go
func TestDog(t *testing.T) {
	//1.go语言不支持显示的类型转换，强转试一下~
	// var p *Pet:=(*Pet)new(Dog)//这里编译器会编译出错，555~ 强转也不行
	dog := new(Dog)
	dog.SpeakTo("Tea")
}
//out:
//...
//Tea
~~~

根据输出结果我们可以看出来，Go语言不支持函数重写，函数重写是无效的，不能向java一样“多态”，要想重写Pet中的方法，我们可以 使用“组合”的形式，将Pet作为Dog的成员，然后在Dog中重写Pet中的方法。

eg:

~~~go
type Dog struct {
	p *Pet
}

func (d *Dog) Speak() {
	fmt.Println("Dog~")
}

func (d *Dog) SpeakTo(host string) {
	d.Speak()
	fmt.Println(" ", host)
}
~~~

从上面的示例中可以看出，这个“多态”其实和java中的多态是不一样的，十分笨重。

### :shark:错误处理

#### 1. error接口

1. 没有异常机制
2. error类型实现了error接口
3. 可以通过error.New来快速创建错误实例

~~~go
type error interface{
    Error() string
}
~~~

eg:

~~~go
func getError() error {
	return errors.New("I am a error ")
}
~~~

#### 2. panic

- panic 用于不可恢复的错误
- panic推出前会执行defer指定的内容

#### os.Exit与panic的区别：

1. os.Exit退出不会调用defer指定的函数
2. os.Exit退出时不会输出当前调用栈信息

```go
func TestPanicVsExit(t *testing.T){

   defer func() {
      fmt.Println("Finally")
   }()

   fmt.Println("Start")

   panic(errors.New("Something wrong !"))

   //os.Exit(-1)

}
```

out:

1. 使用panic会打印堆栈信息,并且会执行defer函数

~~~
=== RUN   TestPanicVsExit
Start
Finally
--- FAIL: TestPanicVsExit (0.00s)
panic: Something wrong ! [recovered]
	panic: Something wrong !

goroutine 20 [running]:
testing.tRunner.func1(0xc0000b0200)
	C:/Go/src/testing/testing.go:830 +0x399
panic(0x521960, 0xc000042530)
	C:/Go/src/runtime/panic.go:522 +0x1c3
command-line-arguments.TestPanicVsExit(0xc0000b0200)
	H:/go-study/src/error/error_test.go:27 +0x107
testing.tRunner(0xc0000b0200, 0x554728)
	C:/Go/src/testing/testing.go:865 +0xc7
created by testing.(*T).Run
	C:/Go/src/testing/testing.go:916 +0x361
~~~

2. 使用os.Exit，程序正常退出，并不会执行defer函数

~~~
=== RUN   TestPanicVsExit
Start
~~~

#### 3. recover错误恢复机制

返回panic传递的异常信息。

```go
func TestRecover(t *testing.T){
   defer func() {
      if err := recover(); err!=nil{
         fmt.Println("recover from ",err)
      }
   }()

   fmt.Println("Start")
   panic(errors.New("Something wrong !"))
}
```

out:

~~~
=== RUN   TestRecover
Start
recover from  Something wrong !
--- PASS: TestRecover (0.00s)
PASS
~~~

### :m:项目管理

#### package

1. 基本复用模块（以首字母大写来表明可被包外代码访问）
2. 代码的package可以和所在的目录保持不一致
3. 同一目录里的Go代码的package要保持一致

##### init

1. 在main函数被执行前，所有依赖的package的init方法都会被执行
2. 不同包的init函数按照导入包的依赖关系决定执行顺序
3. 不同包可以有多个init函数
4. 包的每个源文件也可以有多个init函数，这点比较特殊

#### 调用远程go程序

1. 通过go get 来获取远程依赖
   1. go get -u 强制从网络更新远程依赖
2. 注意代码在Github 上不要有组织形式，以适应go get
   1. 直接以代码路径开始不要有src

#### Go未解决的依赖问题

1. 同意环境下，不同项目使用的同一包的不同保本
2. 无法管理对包的特定版本的依赖

#### verder路径

`Go 1.5release`版本发布，vender目录被添加到GOPATH和GOROOT之外的依赖目录查找解决方案。在Go1.6之前需要手动设置环境变量。

**查找依赖包路径的解决方案如下：**

1. 当前包下的wender目录
2. 向上级目录查找，直到找到src的vender目录
3. 在GOPATH下面查找依赖包
4. 在GOROOT目录下查找

#####  常用的依赖管理工具

- **godep**:https://github.com//tools/godep
- **glide**:https://github.com/Masterminds/glide
- **dep**:https://github.com/golang/dep

## :dog:协程

#### **java线程和协程比较**

1. java创建一个线程默认的Stack size = 1MB

   Groutine的Stack size = 2 K

2. KSE对应关系

   java Thread 和系统线程比是1：1

   Groutine 是M:N

1. java线程机制

![](H:\go-study\images\java线程.png)

java线程和内核线程一一对应，效率很高，但是如果频繁的进行线程的切换也会导致内核线程的切换造成较大的系统消耗。

2. 协程机制：

![](H:\go-study\images\协程.png)

1. 一个系统线程对应一个处理器processor,每个处理器后面跟着一个协程队列，

2. processor依次运行队列中的协程，
3. 在启动程序的时候，会开启一个守护线程用来计算每个processor完成的协程数量。如果一段时间后，某个processor完成的协程数量没有发生变化，守护线程会在协程的任务栈中插入一个特殊的标记，当协程读到这个标记的时候，就会中断自己，并查到该队列的队尾，切换成别的协程进一步继续运行。
4. 当某一个协程被系统中断了，processor会将自己移到另一个可运行的系统线程中，继续执行队列中其他的协程，当中断的协程别唤醒之后，自己会加入到processor等待队列中。当协程被中断的时候，中断现场会被保存到当前协程的对象中，保证协程唤醒后正常继续运行。

#### 编程实现

```go
func TestGroutine(t *testing.T){
   for i:=0;i<10;i++{
      go func(i int) {
         fmt.Println(i)
      }(i)
   }
}
```

out:

~~~
=== RUN   TestGroutine
2
1
9
3
4
5
6
7
8
0
--- PASS: TestGroutine (0.00s)
~~~

### 并发控制

```go
func TestCounter(t *testing.T) {
   counter := 0
   for i := 0; i < 5000; i++ {
      go func() {
         counter++
      }()
   }
   time.Sleep(1 * time.Second)
   t.Logf("counter = %d", counter)
}
```

out:

~~~
groutine_test.go:25: counter = 4712
~~~

上一个例子中，系统会开启5000个协程，完成对counter计数，因为没用考虑并发安全机制，导致计算的结果和预想的不一致。

**解决方案 **

使用sync.Mutex{} 的Lock()和UnLock()方法

```
func TestCounterThreadSafe(t *testing.T) {
   mut:=sync.Mutex{}
   counter := 0
   for i := 0; i < 5000; i++ {
      go func() {
         defer func() {
            mut.Unlock()
         }()
         mut.Lock()
         counter++
      }()
   }
   //让外面的代码执行的速度慢一点，因为如果外面的代码执行完后，可能协程还没有执行完，导致
   //结果和预想的不一致，程序出现问题
   time.Sleep(1 * time.Second)
   t.Logf("counter = %d", counter)
}
```

看着加锁和释放锁的部分可能有点难受~~~，我也难受，但是go语言的机制确实可以这样写:happy:

#### WaitGroup

解决外部执行时间和协程执行时间不一致的问题

```go
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
```

### CSP并发机制

Communicating sequential processes

依赖通道完成两个通信实体之间的通信。

CSP vs Actor

- 和Actor的直接通讯不同，CSP模式则是通过Channel进行通讯的，更松耦合一些
- Go中的channel是有容量限制的并且独立于处理Groutine,而如Erlang,Actor模式中的mailbox容量是无限的，接收进程也总是被动的处理消息。

#### Channel

同步channel，会阻塞等待结果的协程。

异步channel，不会阻塞等待结果的协程。

**声明方式：**

同步：`retch:=make(chan string)`声明一个string类型的同步channel

异步：`retch:=make(chan string,1)`声明一个string类型的异步channel，并指定容量大小为1

































