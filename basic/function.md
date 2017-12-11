函数function
=================================

	* Go语言函数不支持嵌套、重载和默认参数
	* 但支持一下特性：
        无需声明函数原型、不定长度变参、多返回值、命名返回值参数、匿名函数、闭包

	* 定义函数使用关键字func，且左大括号不能另起一行
	* 函数也可以作为一种类型使用
	* Go语言使用的是值传递，即在调用过程中不会影响到实际参数


**函数定义方式**
```go
package main
import (
    "fmt"
)
func main() {
    f6(5)
}
//无参无返回值
func f1() {
}
//参数列表
func f2(a int, b string) {
}
//相同类型参数列表
func f3(a, b, c int) {
}
//带返回值，如果只返回一个参数，则可以省略小括号
func f4() (int, string, int) {
    return 0, "str", 1
}
//指定返回参数，返回值类型相同时，可省略类型，写成(x,y,z int)
func f5() (x string, y int, z int) {
    return
}
//不定长变参，不定长变参必须是最后一个参数，也就是不定长变参后面不能再跟参数
func f6(a ...int) {
    //这里a是一个slice
    fmt.Println(a)
}
```


    Go语言中使用值传递，函数调用时参数都是值拷贝。
```go
func main() {
    a := 1
    f(a)           //输出2
    fmt.Println(a) //输出1，函数f中对a的操作并不会影响这里的a
}
func f(x int) {
    x = 2
    fmt.Println(x)
}
/*
输出结果：
    2
    1
*/
``` 
slice比较特殊，函数直接以slice作为参数，调用时传递的是slice的内存地址拷贝，所以在函数中对slice进行的操作，是直接作用到源slice上的。如下，
```go
package main
import (
    "fmt"
)
func main() {
    s := []int{1, 2, 3, 4}
    f(s)     //这里传递的是slice的内存地址拷贝
    fmt.Println(s)
}
func f(s []int) {
    s[0] = 5
    s[1] = 6
    s[2] = 7
    s[3] = 8
    fmt.Println(s)
}
/*
输出结果：
    [5 6 7 8]
    [5 6 7 8]
*/
```

使用指针来传递，可以通过调用函数来修改变量的值。如下，
```go
func main() {
    a := 1
    f(&a)    //传递a的地址值
    fmt.Println(a)
}

func f(x *int) { 
    *x = 2
    fmt.Println(*x)
}
/*
输出结果：
    2
    2
*/
```

函数作为类型

```go
package main
import (
    "fmt"
)
func main() {
    a := f
    a()     //输出：Func f
}
func f() {
    fmt.Println("Func f")
}
```


匿名函数

```go
package main
import (
    "fmt"
    "os"
)
func no_func() {

    // 匿名函数 1
    f := func(i, j int) (result int) {     // f 为函数地址
        result = i + j
        return result
    }
    fmt.Fprintf(os.Stdout, "f = %v  f(1,3) = %v\n", f, f(1, 3))

    // 匿名函数 2
    func(i, j int) (m, n int) {     // x y 为函数返回值
        return j, i
    }(1, 9)     // 直接创建匿名函数并执行
}

func main() {
    //调用
    no_func()
}
```

闭包
```go
package main
import (
    "fmt"
)

func main() {
    f := closure(10)

    fmt.Println(f(1))     //输出：11
    fmt.Println(f(2))     //输出：12
}

func closure(x int) func(int) int {
    fmt.Printf("%p\n", &x)
    return func(y int) int {
        fmt.Printf("%p\n", &x)
        return x + y
    }
}
/*
输出结果：
    0xc042048080
    0xc042048080
    11
    0xc042048080
    12
*/
```

三次取的x的地址值是相同的，说明三次调用的x都是同一个，而不是新的拷贝。


Go语言中defer关键字
-------------------------

    * defer的执行方式类似其他语言中的析构函数，在函数体执行结束后按照调用顺序的相反顺序逐个执行
    * 即使函数发生眼中错误也会执行
    * 支持匿名函数的调用
    * 常用于资源清理、文件关闭、解锁以及记录时间等操作
    * 通过与匿名函数配合可在return之后修改函数计算结果
    * 如果函数体内某个变量作为defer时匿名函数的参数，则在定义defer时即已经获得了拷贝，否则则是引用某个变量的地址



    * Go没有异常机制，但有panic/recover模式来处理错误
    * panic可以在任何地方引发，但recover只有在defer调用的函数中有效

