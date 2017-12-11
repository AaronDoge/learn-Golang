Go语言中的结构Struct
========================

	* Go中的struct与C中的struct非常相似，并且Go没有class
	* 使用type<Name> struct{}定义结构，名称遵循可见性规则
	* 支持指向自身的指针类型成员
	* 支持匿名结构也可以用于map的值value
	* 可以使用字面值对结构进行初始化
	* 允许直接通过指针来读写结构成员
	* 相同类型的成员可进行直接拷贝赋值
	* 支持==与!=比较运算符，但不支持>或<
	* 支持匿名字段，本质上是定义了以某个类型名为名称的字段
	* 嵌入结构作为匿名字段看起来像继承，但不是继承
	* 可以使用匿名字段指针


结构struct定义
--------------------------
```go
package main
import (
    "fmt"
)
type person struct {
    Name string
    Age  int
}
func main() {
    a := person{}
    a.Name = "Joe"
    a.Age = 19
    fmt.Println(a)
}
```

或者如下定义（注意，逗号！！！不要丢）
```go
func main() {
    a := person{
        Name: "Joe",
        Age:  19,        //注意，逗号！！！
    }
    // a.Name = "Joe"
    // a.Age = 19
    fmt.Println(a)
}
```
>**注意：**逗号千万不要丢。

结构struct作为函数参数时传递的是struct的值拷贝
```go
package main
import (
    "fmt"
)
type person struct {
    Name string
    Age  int
}
func main() {
    a := person{
        Name: "Joe",
        Age:  19,
    }
    // a.Name = "Joe"
    // a.Age = 19
    fmt.Println(a)    //{Joe 19}
    f(a)                      //f {Joe 17}
    fmt.Println(a)    //{Joe 19}，并没有因为调用函数f而改变，可见函数f传递的参数是值拷贝
}
func f(per person) {  //可以看出，struct是作为一种类型
    per.Age = 17
    fmt.Println("f", per)
}
/*
输出结果：
    {Joe 19}
    f {Joe 17}
    {Joe 19}
*/
```
可以使用指针
```go
type person struct {
    Name string
    Age  int
}
func main() {
    a := person{
        Name: "Joe",
        Age:  19,
    }
    // a.Name = "Joe"
    // a.Age = 19
    fmt.Println(a)     //{Joe 19}
    f(&a)                    //f {Joe 17}
    fmt.Println(a)     //{Joe 17}，调用函数f后，结构体的数据被改变。
}
func f(per *person) {
    per.Age = 17
    fmt.Println("f", *per)
}
/*
输出结果：
    {Joe 19}
    f {Joe 17}
    {Joe 17}
*/
```

或者在结构初始化的时候直接
```go
type person struct {
    Name string
    Age  int
}
func main() {
    a := &person{     //推荐这种方式
        Name: "Joe",
        Age:  19,
    }
    a.Name = "Aaron"     //虽然a是取的person的地址，但还是可以使用这种方式操作struct中的数据的
    // a.Age = 19

    fmt.Println(*a)     //{Aaron 19}
    f(a)                        //f {Aaron 17}
    fmt.Println(a)      //&{Aaron 17}
}

func f(per *person) {
    per.Age = 17
    fmt.Println("f", *per)
}
/*
输出结果：
    {Aaron 19}
    f {Aaron 17}
    &{Aaron 17}
*/
```

匿名struct
-------------------------
```go
package main
import (
    "fmt"
)

func main() {
    a := &struct {     //&
        Name string
        Age  int
    }{
        Name: "Joe",
        Age:  19,        //逗号不能丢！！否则会报错！！
    }
    fmt.Println(a)     //输出：&{Joe 19}
}
```

嵌套struct
------------------------
```go
package main
import (
    "fmt"
)
type person struct {
    Name    string
    Age     int
    Contact struct {
        //Phone, City string //可以写成这种形式，从而省略一个string
        Phone string
        City  string
    }
}
func main() {
    a := &person{
        Name: "Joe",
        Age:  19,
    }
    a.Contact.Phone = "1381234567"
    a.Contact.City = "Shanghai"     //Contact里面的字段只能通过这种方式赋值
    fmt.Println(*a)     //输出：&{Joe 19}
}
```


struct匿名字段（也称嵌入字段）
-------------------------------------

Go语言中，定义一个struct时字段名与其类型一一对应，实际上Go语言支持只提供类型，而不写字段名的方式，也就是匿名字段或嵌入字段。
初始化时必须按照对应的顺序进行，否则会报错。
```go 
type person struct {
    string
    int
}
func main() {
    a := &person{"Joe", 19}
    fmt.Println(*a) //输出：{Joe 19}
}
```

字段的继承
--------------------

Go支持只提供类型而不写字段名的方式，也就是匿名字段，也称嵌入字段。如下面例子中，Student访问属性age，name以及weight的时候，就像访问自己所有用的字段一样，匿名字段能够实现字段的继承。
当匿名字段是一个struct的时候，那么struct所拥有的全部字段都被隐式地引入到当前定义的struct。
下例中，human是struct，本身也是一种类型，所以在teacher/student结构中属于匿名字段。
```go
package main
import (
    "fmt"
)
type human struct {
    Sex string
}
type teacher struct {
    human     //匿名字段，human是一个struct，本身就是一个类型。这里类似于继承，将human结构体嵌入到teacher
    Name  string
    Age   int
}
type student struct {
    human
    Name string
    Age  int
}
func main() {
    a := teacher{Name: "Joe", Age: 19, human: human{Sex: "女"}} //初始化结构
    b := student{Name: "Aaron", Age: 13, human: human{Sex: "男"}}
    //调用结构中的数据
    a.Name = "Julia"
    a.Age = 20
    a.Sex = "不明" //这里可以接通过a.Sex访问到嵌入结构的数据Sex
    fmt.Println(a, b) //输出；    {{不明} Julia 20} {{男} Aaron 13}
}
```

再看一个例子，
```go
package main
import (
    "fmt"
)
type Skills []string
type Human struct {
    name   string
    age    int
    weight int
}
type Student struct {
    Human      //struct类型作为匿名字段
    Skills         //自定义的类型string slice作为匿名字段
    int             //内置类型作为匿名字段
    //int          //也就是说匿名字段，不能有重复类型；如果是重复则必须写出字段名
    string       //相同的类型只能有一个匿名字段，其他的必须有字段名
    speciality string
}
func main() {
    //初始化学生Jane
    Jane := Student{Human: Human{"Jane", 20, 60}, speciality: "CS"}
    //访问字段
    fmt.Println("Her name is", Jane.name)
    fmt.Println("Her speciality is", Jane.speciality)
    //修改字段
    Jane.Skills = []string{"Java", "Golang"}
    fmt.Println("Her skills are", Jane.Skills)

    //修改匿名内置类型字段，内置类型匿名字段的赋值方法
    Jane.int = 3
    Jane.string = "str"
    fmt.Println(Jane.int)
    fmt.Println(Jane.string)
}
/*
输出结果：
    Her name is Jane
    Her speciality is CS
    Her skills are [Java Golang]
    3
    str
*/
```

**同级冲突**
```go
type A struct {
    B
    C
    Name string
}
type B struct {
    Name string
}
type C struct {
    Name string
}
func main() {
    a := A{Name: "Name_A", B: B{Name: "Name_B"}}
    a := A{B: B{Name: "Name_B"}, C: C{Name: "Name_C"}} //报错，B和C在A中属同一级，因此二者的Name相冲突
    fmt.Println(a.Name, a.B.Name)                      //输出：    Name_A Name_B
}
```
