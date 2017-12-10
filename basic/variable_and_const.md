变量的声明和赋值
===========================

单个变量的声明和赋值
---------------------------

变量的声明格式：`var <变量名称> <变量类型>`
变量的赋值格式：`<变量名称> = <表达式>`
声明的同时赋值：`var <变量名称> [变量类型] = <表达式>`
```go
var a int     //变量的声明
a = 123        //变量的赋值

//变量声明的同时赋值（一般用在全局位置）
var b int = 123
//上行的格式可省略变量类型，由系统推断
var c = 123
//变量声明与赋值的最简单写法（全局变量位置不能使用）
d := 345
```

多个变量的声明和赋值
----------------------------

	* 全局变量的声明可以使用var()的方式进行简写
	* 全局变量的声明不可以省略var，但可以使用并行方式
	* 所有变量都可以使用类型推断
	* 局部变量不可以使用var()的方式简写，只能使用并行方式
	* 局部变量可以省略var，全局变量不可以
```go
//全局位置
var(
    //使用常规方式
    a = "hello"
    //使用并行方式以及类型推断
    b, s, d = 1, 2, 3
    //c := 4 //不可以省略var，不可以使用:=的方式
)
```


Go中有一个空白符`“ _ ”`，表示空位:`a, _, c, d := 1, 2, 3, 4`


变量的类型转换
-------------------------------

	* Go中不存在隐式转换，所有类型转换必须显示声明
	* 转换只能发生在两种相互兼容的类型之间
	* 类型转换的格式：<ValueA> [:]= <TypeOfValueA>(<ValueB>)


因此，Go语言是类型安全的。

```go
//在相互兼容的两种类型间进行转换
var a float32 = 1.1
b := int(a)

//以下无法通过编译
var c bool = true
d := int(c)
```

例子，
```go
func main() {
    var a int = 65
    b := string(a)
    fmt.Println(a)
    fmt.Println(b)
}

//输出
//65
//A
```
`string()`表示将数据转换成文本格式，因为计算机中存储的任何东西本质上都是数字，因此此函数自然地认为我们需要的是用数组65表示的文本A。
如果需要字符串转换为"65"，则需先导入包`strconv`，然后调用`Itoa()`方法进行转换。反过来转换则可以调用`Atoi()`方法。
```go
import (
    "fmt"
    "strconv"
)
type (
    byte int8
    rune int32
)
func main() {
    var a int = 65
    b := string(a)
    c := strconv.Itoa(a)
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
}
```

常量的定义、初始化和枚举
================================

常量的定义
--------------------------------

	* 常量的值在编译时就已经确定
	* 常量的定义格式与变量基本相同
	* 等号右侧必须是常量或者常量表达式
	* 常量表达式中的函数必须是内置函数（内置函数在编译时就已经确定）

```go
package main
import (
    "fmt"
)
const a int = 1
const b = 'A'
const f, g, h = 3, "ni", 'N'
const (
    c = 5
    d
    e
    //d和e会被初始化为c的值，即d=5,e=5
)
func main() {
    fmt.Println(a) //输出1
    fmt.Println(b) //输出65
    fmt.Println(c) //输出5
    fmt.Println(d) //输出5
    fmt.Println(e) //输出5
    fmt.Println(f) //输出3
    fmt.Println(g) //输出ni
    fmt.Println(h) //输出78
}
```
 **几个需要注意的细节:**
1.
```go
var ss = "abc"

const (
    a = len(ss) //编译报错：const initializer must be constant
    b
    c
)
```
注意，常量的值必须在编译时就确定，这里全局变量在编译时并没有被处理，len(ss)不能被赋值给常量a。
```go
const (
    a = "123"
    b = len(a) //这个是可以的
)
```
2.
```go
const (
    a, b = 1, 2
    c
)
//这里编译会报错，c不能确定被赋于a的值，还是b的值。

const (
    a, b = 1, 2
    c, d
)
//这样写是可以的，此时c=1，d=2
```

常量的初始化规则与枚举
--------------------------------

	* 在定义常量组时，如果不提供初始值，则表示将使用上一行的表达式
	* 使用相同的表达式不代表具有相同的值
	* iota是常量的计数器，从0开始，组中每定义一个常量自动递增1
	* 通过初始化规则与iota可以达到枚举的效果
	* 每遇到一个const关键字，iota就会重置为0

```go
package main
import (
    "fmt"
)
const (
    a = 'A'
    b
    c = iota
    d
)
const (
    //第一个常量不可省略表达式
    Monday = iota
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
    Sunday
)
func main() {
    fmt.Println(a) //输出65
    fmt.Println(b) //输出65
    fmt.Println(c) //2
    fmt.Println(d) ////3，注意这里输出的是3，表明d=iota，而不是等于c的值2

    fmt.Println(Monday)    //0
    fmt.Println(Tuesday)   //1
    fmt.Println(Wednesday) //2
    fmt.Println(Thursday)  //3
    fmt.Println(Friday)    //4
    fmt.Println(Saturday)  //5
    fmt.Println(Sunday)    //6
}
```

计数器`iota`，每遇到`const`都会重置为0.
