Go语言条件语句
===========================

判断语句if
---------------------------

	* 条件表达式没有括号
	* 支持一个初始化表达式（可以是并行方式）
	* 左大括号和条件语句或else必须在同一行
	* 支持单行模式
	* 初始化语句中的block级别，同时隐藏外部同名变量

```go
package main
import (
    "fmt"
)
func main() {
    var a int = 10
    //不带括号
    if a > 3 { //这个括号必须和条件语句或else同一行
        fmt.Println(a)
    }
    if a, b := 1, 2; a > 0 && b > 1 {
        fmt.Print("支持初始化表达式，a的值为")
        fmt.Println(a)
    }
    //这里a的作用范围仅限于if语句，而且对于像这样的内外同名变量，内部的变量会覆盖外部的变量。
    fmt.Println(a)
}
```

GO语言循环语句for
--------------------------------

	* Go只有for一个循环语句关键字，但支持3中形式
	* 初始化和步进表达式可以是多个值
	* 条件语句每次循环都会被重新检查，因此不建议在条件语句中使用函数，尽量提前计算好条件并以变量或常量代替
	* 左大括号必须和条件语句在同一行

```go
//for的三种形式
//第一种
func mian(){
    a := 1
    for {
        a++
        if a > 3 {
            break
        }
    }
    fmt.Println(a)
}
//第二种
func main(){
    a := 1
    for a <= 3 {
        a++
    }
    fmt.Println(a)
}
//第三种
func main(){
    a := 1
    for i := 0; i < 3; i++ {
        a++
    }
    fmt.Println(a)
}
```

选择语句switch
-------------------------------------

	* 可以使用任何类型或表达式作为条件语句
	* 不需要写break，一旦条件符合会自动终止
	* 如希望继续执行下一个case，需要使用fallthrough语句
	* 支持一个初始化表达式（可以是并行方式），右侧需要跟分号
	* 做大括号必须和条件语句在同一行


```go
func main(){
    a := 1
    switch a {
        case 0:
            fmt.Println("a=0")
            //这里不需要break，自动终止。如果想让程序继续，可以使用fallthrough语句
        case 1:
            fmt.Println("a=1")
        default:
            fmt.Println("None")
    }
    fmt.Println(a)
}

//
func main(){
    a := 10
    switch a := 1; { //不要忘记这里有个“;”
        case a > 0:
            fmt.Println("a>0")
            fallthrough
        case a >= 1:
            fmt.Println("a>=1")
        default:
            fmt.Println("None")
    }
    fmt.Println(a)    //输出10
}
```





跳转语句goto/break/continue
-----------------------------------------

	* 三个语法都可以配合标签使用
	* 标签名区分大小写，若标签创建了却不使用会造成编译错误：`label LABEL defined and not used`
	* break和continue配合标签可用于多层循环的跳出
	* goto是调整执行位置，与其他两个语句配合标签的结构并不相同


标签配合goto使用时，标签放在goto前面容易造成死循环。
