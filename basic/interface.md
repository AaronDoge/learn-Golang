
Go语言中的接口interface
=============

接口是一个或多个方法签名（没有方法体）的组合，通过interface来定义对象的一组行为

interface类型定义了一组方法，如果某个类型对象实现了该接口的所有方法签名，即算实现该接口，无需显式实现该接口。

> 注意！实现接口中的方法时，receiver不能是指针。

```go
package main
import (
    "fmt"
)
type Human struct {
    name  string
    age   int
    phone string
}
type Student struct {
    Human
    school string
    loan   float32
}
type Employee struct {
    Human
    company string
    money   float32
}

//Human对象实现SayHi方法，  注意！receiver不能是指针  
func (h Human) SayHi() {
    fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}
//Human对象实现Sing方法
func (h Human) Sing(lyrics string) {
    fmt.Println("lala, la la la la la la...", lyrics)
}
//Human对象实现Guzzle方法
func (h Human) Guzzle(beerStein string) {
    fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}
//Employee重载Human的SayHi方法
func (e Employee) SayHi() {
    fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}
//Student实现BorrowMoney方法
func (s Student) BorrowMoney(amount float32) {
    s.loan += amount
}
//Employee实现SpendSalary方法
func (e Employee) SpendSalary(amonut float32) {
    e.money -= amonut
}
func main() {
    mike := Student{Human{"Mike", 25, "111-111111"}, "MIT", 0.00}
    // paul := Student{Human{"Paul", 26, "111-222211"}, "Harvard", 100}
    // sam := Employee{Human{"Sam", 32, "333-3232124"}, "Google Inc.", 1000}
    tom := Employee{Human{"Tom", 38, "444-12432345"}, "Amazon", 1500}

    //定义Man的变量i
    var i Men
    //i能存储Student，有点类似多态
    i = mike
    fmt.Printf("This is %s, a student.\n", mike.name) //但是这里并不能使用i.name访问到Mike的name字段，因为i是Man类型的，Man中并没有字段
    i.SayHi()
    i.Sing("November rain")

    //i也能存储Employee
    i = tom
    fmt.Printf("This is %s, a Employee.\n", tom.name)
    i.SayHi()
    i.Sing("Born to be wild")
}
type Men interface {
    SayHi()
    Sing(lyrics string)
    Guzzle(beerStein string)
}
type YoungChap interface {
    SayHi()
    Sing(song string)
    BorrowMoney(amount float32)
}
type ElderlyGent interface {
    SayHi()
    Sing(song string)
    SpendSalary(amount float32)
}
```

一个interface的变量，可以存实现这个interface的任意类型对象。例如上面的例子中，定义了一个Men的interface类型的变量i，那么可以给i赋值Human、Student或者Employee等。
因为i能够持有这三种类型的对象，所以我们可以定义一个包含Men类型元素的slice，这个slice可以被赋予实现了Men接口的任意结构的对象，这个和我们传统意义上面的slice有所不同。


空interface
----------------

空接口不包含任何的method，正因如此，所有的类型都实现了空interface。空接口对于描述起不到任何的作用，但是空接口在我们需要存储任意类型的数值的时候相当有用，因为它可以存储任意类型的数值。它优点类似于C语言void\*类型。
```go
func main() {
    //接口还可以这样定义。
    var a interface{}
    var i int = 5
    s := "hello"

    //a可以存储任意类型的数值
    a = i
    fmt.Println(a)
    a = s
    fmt.Println(a)
}
```
一个函数把interface{}作为参数，那么他可以接收任意类型的值作为参数，如果一个函数返回interface{}，那么也就可以返回任意类型的的值。



interface函数参数
-------------------

interface的变量可以持有任意实现该interface类型的对象，这给我们编写函数（包括method）提供了一些额外的思考，我们可以通过定义interface参数，让函数接受各种类型的参数。

例如，fmt.Println是我们常用的一个函数，他可以接收任意数据类型。看fmt的源码文件，有如下代码：
```go
type Stringer interface{ 
    String() string
}
```

也就是说，任何实现了String() string 方法的类型都能作为参数被fmt.Println调用。看如下代码：
```go
package main
import (
    "fmt"
    "strconv"
)

type Human struct {
    name  string
    age   int
    phone string
}

//通过这个方法Human实现了fmt.Stringer接口
func (h Human) String() string {
    return "（" + h.name + " - " + strconv.Itoa(h.age) + " years - Phone:" + h.phone + "）"
}
func main() {
    Bob := Human{"Bob", 35, "000-7777732"}
    fmt.Println("This Human is:", Bob)     //Println打印的是String()方法的返回值
}
```

如果某个类型需要被fmt包以特殊格式输出，那么这个类型就必须实现Stringer接口，否则fmt将以默认的方式输出。


interface变量存储的类型
-------------------------
    
interface的变量里面可以存储任意类型（实现了该interface的类型）的数值。如何反向知道这个变量里面实际保存了的是哪个类型的对象呢？目前有两种方法：

Comma-ok断言
-------------------------
Go语言中可通过如下语法判断是否是该类型的变量：value, ok := element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。
    如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false。
```go
package main
import (
    "fmt"
    "strconv"
)
type Element interface{}
type List []Element //这里List其实是定义了一个类型别名
type Person struct {
    name string
    age  int
}
//定义String方法，实现了fmt.Stringer接口
func (p Person) String() string {
    return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + "years)"
}

func main() {
    list := make(List, 3)
    list[0] = 1
    list[1] = "Hello"
    list[2] = Person{"Dennis", 20}

    for index, element := range list {
        if value, ok := element.(int); ok {        //判断value是否是int类型，是ok=true，否则ok=false
            fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
        } else if value, ok := element.(string); ok {
            fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
        } else if value, ok := element.(Person); ok {
            fmt.Printf("list[%d] is a Person and its vlaue is %s\n", index, value)
        } else {
            fmt.Printf("list[%d] is of a different type", index)
        }
    }
}
/*
输出结果：
    list[0] is an int and its value is 1
    list[1] is a string and its value is Hello
    list[2] is a Person and its vlaue is (name: Dennis - age: 20years)
*/
```

上面例子中断言的类型越多，那么if else也就越多，所以引出下面switch

switch测试
----------------------
```go  
package main
import (
    "fmt"
    "strconv"
)

type Element interface{}

type List []Element

type Person struct {
    name string
    age  int
}
func (p Person) String() string {
    return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func main() {
    list := make(List, 3)
    list[0] = 1
    list[1] = "Hello"
    list[2] = Person{"Dennis", 23}

    for index, element := range list {
        switch value := element.(type) {
        case int:
            fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
        case string:
            fmt.Printf("list[%d] is an string and its value is %s\n", index, value)
        case Person:
            fmt.Printf("list[%d] is an Person and its value is %s\n", index, value)
        default:
            fmt.Printf("list[%d] is of a different type", index)
        }
    }
}
/*
输出结果：
    list[0] is an int and its value is 1
    list[1] is an string and its value is Hello
    list[2] is an Person and its value is (name: Dennis - age: 23 years)
*/
```



嵌入interface
--------------------------

如果一个interface1作为interface2的一个嵌入字段，那么interface2隐式的包含了interface1里面的methods。


反射
----------------

Go语言实现了反射，所谓反射就是能检查程序在运行时的状态。


	* 反射可大大提高程序的灵活性，是的interface{}有更大的发挥余地
	* 反射使用TypeOf和ValueOf函数从接口中获取目标对象信息
	* 反射会将匿名字段作为独立字段（匿名字段本质）
	* 想要利用反射修改对象状态，前提是interface.data，是settable，即pointer-interface
	* 通过反射可以动态调用方法

```go
package main
import (
    "fmt"
    "reflect"
)

type User struct {
    Id   int
    Name string
    Age  int
}

func (u User) Hello() {
    fmt.Println("Hello Wrold.")
}

func main() {
    u := User{123, "Joe", 12}
    Info(u)
}

func Info(o interface{}) { //接口类型作为函数参数，这里是空接口，意味着可以接受任意类型。
    //类型
    t := reflect.TypeOf(o)        //返回作为参数传进来的对象(该对象实现了上行的接口参数)，这里是User
    fmt.Println(t)                //main.User
    fmt.Println("Type", t.Name()) //Type User，获取类型名称

    //当传入参数为指针时，是不能正常使用下面的反射的功能的，可以加如下判断
    if k := t.Kind(); k != reflect.Struct {
        fmt.Println("输入的类型不符合要求...")
        return
    }

    fmt.Println(t.NumField())    //3，传入对象o(这里是User)的字段个数
    fmt.Println(t.Field(0).Name) //Id，获取第0个字段名称


    //值
    v := reflect.ValueOf(o)
    fmt.Println(v)                      //{1 Joe 12}，v接受User的value
    fmt.Println(v.Field(0))             //123，获取第0个字段的值
    fmt.Println(v.Field(0).Interface()) //123，???

    fmt.Println("Fields:")              //循环获取每个字段的值
    for i := 0; i < t.NumField(); i++ {
        f := t.Field(i)
        val := v.Field(i).Interface()
        fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
    }


    //获取User的方法
    for i := 0; i < t.NumMethod(); i++ {
        m := t.Method(i)
        fmt.Printf("%6s: %v\n", m.Name, m.Type)
    }
}
```


使用反射修改数据
```go
func main() {
    x := 123
    v := reflect.ValueOf(&x)
    v.Elem().SetInt(999)

    fmt.Println(x)     //999
}

package main
import (
    "fmt"
    "reflect"
)

type User struct {
    Id   int
    Name string
    Age  int
}

func main() {
    u := User{123, "Joe", 12}

    Set(&u)     //这里传入指针，t.Kind()为ptr；直接传入u，则t.Kind()是struct
    fmt.Println(u)
}

func Set(o interface{}) {
    v := reflect.ValueOf(o)

    if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
        fmt.Println("xxx")
        return
    } else {
        v = v.Elem()
    }

    f := v.FieldByName("Name")
    //判断f是否获取成功
    if !f.IsValid() {
        fmt.Println("Bad")
        return
    }

    if f.Kind() == reflect.String {
        f.SetString("ByeBye")     //对Name属性进行修改
    }
}
/*
输出结果：
    {123 ByeBye 12}
*/
```

调用带参方法
```go
package main
import (
    "fmt"
    "reflect"
)

type User struct {
    Id   int
    Name string
    Age  int
}

func (u User) Hello(name string, times int) {
    fmt.Println("Hello", name, ", my name is ", u.Name, "and you can call me at", times)
}

func main() {
    u := User{1, "Joe", 23}

    //u.Hello("everyone")
    v := reflect.ValueOf(u)

    //通过v调用MethodByName函数调用方法
    mv := v.MethodByName("Hello")
    args := []reflect.Value{reflect.ValueOf("Jhon"), reflect.ValueOf(15)}     //把参数按顺序放入slice
    mv.Call(args)                                                         //mv通过Call函数调用方法，并传入参数args
}

```
