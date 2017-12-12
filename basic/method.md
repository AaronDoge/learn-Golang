Go语言方法method
=================================

	* Go中虽然没有class，但依旧有method
	* 通过显式说明receiver来实现与某个类型的结合
	* 只能为同一个包中的类型定义方法
	* Receiver可以是类型的值或者指针
	* 不存在方法重载
	* 可以使用值或者指针来调用方法，编译器会自动完成转换
	* 从某种意义上来说，方法是函数的语法几乎一样，因为receiver其实就是方法所接收的第一个参数（Method Value ，Method Expression）
	* 如果外部结构和嵌入结构存在同名方法，则优先调用外部结构的方法
	* 类型别名不会拥有底层类型所附带的方法
	* 方法可以调用结构中的非公开字段



考虑求几何图形的面积问题这一情形，定义一种几何图形的struct，以长方形为例，按照一般思路，会考虑使用函数来实现，如下
```go
package main
import (
    "fmt"
    "math"
)

type Rectangle struct {
    width, height float64
}

type Circle struct {
    radius float64
}

func area(r Rectangle) float64 {
    return r.height * r.width
}

// func area(c Circle) float64 {     //Go语言中并不能重载函数，所以这样写是不对的
//     return c.radius * c.radius * math.Pi
// }

func main() {
    r1 := Rectangle{12, 2}
    r2 := Rectangle{9, 4}
    fmt.Println("Area of r1 is : ", area(r1))
    fmt.Println("Area of r2 is : ", area(r2))
}
```

Go语言中并不支持函数重载，当需要增加圆形、五边形等其他几何图形的时候只能增加新的函数`area_circle`，`area_triangle`等。这样的实现在调用的时候就会比较麻烦。
其实，Go语言中也有类似函数重载的机制，这就是Go中的method的概念了。method是附属在一个给定类型上的，它的语法和函数的声明语法几乎一样，只是在func后面增加了一个`receiver`（也就是method所依从的主体）。

用上面提到的形状的例子来说，`method area()`是依赖于某个形状（比如说Rectangle）来发生作用的。Rectangle.area()的发出者是Rectangle，area()是属于Rectangle的方法，而非一个外围函数。
更具体的说，Rectangle存在字段length和width，同时存在方法area()，这些字段和方法都属于Rectangle。

上面的例子用method来实现，如下
```go
package main
import (
    "fmt"
    "math"
)

type Rectangle struct {
    width, height float64
}
type Circle struct {
    radius float64
}

//定义求Rectangle面积的方法method
func (r Rectangle) area() float64 {
    return r.width * r.height
}
//求Circle面积的方法method。这个方法和上面方法名相同，但属于不同结构。类似于Java中的方法重载。
func (c Circle) area() float64 {
    return c.radius * c.radius * math.Pi
}

func main() {
    r1 := Rectangle{12, 2}
    r2 := Rectangle{9, 4}
    c1 := Circle{3}
    c2 := Circle{5}
    fmt.Println("Area of r1 is : ", r1.area()) //调用方法
    fmt.Println("Area of r1 is : ", r2.area())
    fmt.Println("Area of r2 is : ", c1.area())
    fmt.Println("Area of r1 is : ", c2.area())
}
/*
输出结果：
    Area of r1 is :  24
    Area of r1 is :  36
    Area of r2 is :  28.274333882308138
    Area of r1 is :  78.53981633974483
*/
```

> 使用method需要注意的几点：
> 1. 虽然method的名字一样，但是如果接收者receiver不一样，method就不一样
> 2. method里面可以访问接收者的字段
> 3. 调用method通过“.”访问，就像struct里面访问字段一样
> 4. 方法的receiver是值传递，而非引用传递。Receiver还可以是指针，二者的差别在于，指针作为receiver会对实例对象的内容发生操作，而普通类型作为receiver仅仅是以副本作为操作对象，并不会对原实例对象发生操作。

```go
package main
import (
    "fmt"
)

const (
    WHITE = iota
    BLACK
    BLUE
    RED
    YELLOW
)

type Color byte    //类型别名

type Box struct {
    width, height, depth float64
    color                Color
}

type BoxList []Box     //a slice of boxes

//Volume方法的接受者为Box，并返回Box的容量
func (b Box) Volume() float64 {
    return b.width * b.height * b.depth
}

//SetColor方法的接收者为指向Box的指针，通过指针对Box的Color进行修改
func (b *Box) SetColor(c Color) {
    b.color = c     //虽然b是指针，但可以直接用b.color的方式访问color
}

//BiggestColor方法的接收者为BoxList，BoxList是一个slice类型（而不是struct了），并返回Color类型数据
func (bl BoxList) BiggestColor() Color {
    v := 0.00
    k := Color(WHITE)     //？？？这种初始化方式，a:=int(3)，b:=byte(RED)
    for _, b := range bl {
        if bv := b.Volume(); bv > v {
            v = bv
            k = b.color
        }
    }
    return k
}

//BoxList为接收者，BoxList是一个slice，这里考虑slice作为参数的特殊性。
func (bl BoxList) PaintItBlack() {
    for i, _ := range bl {
        bl[i].SetColor(BLACK)    //这里对slice的修改即是对源slice的修改
    }
}

//接收者是Color也就是byte类型，并返回string类型
func (c Color) String() string {
    strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
    return strings[c]
}

func main() {
    boxes := BoxList{
        Box{4, 4, 4, RED},
        Box{1, 1, 20, BLACK},
        Box{10, 10, 1, BLUE},
        Box{10, 30, 1, WHITE},
        Box{20, 20, 20, YELLOW},
    }

    fmt.Printf("We have %d boxes in our set\n", len(boxes))
    fmt.Println("The volume of the first one is", boxes[0].Volume(), "立方厘米")
    fmt.Println("The color of the last one is", boxes[len(boxes)-3].color.String())
    fmt.Println("The biggest one is", boxes.BiggestColor().String())
    fmt.Println("Let's Paint them all black")
    boxes.PaintItBlack()
    fmt.Println("The color of the second one is", boxes[1].color.String())
    fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())
}
/*
输出结果：
    We have 5 boxes in our set
    The volume of the first one is 64 立方厘米
    The color of the last one is BLUE
    The biggest one is YELLOW
    Let's Paint them all black
    The color of the second one is BLACK
    Obviously, now, the biggest one is BLACK
*/
```

指针作为receiver
----------------------

struct作为接收者receiver，传递到方法中的是struct的一个值拷贝，这时在方法中对struct的修改都是对拷贝副本的修改，而不会影响到源struct。如果想对原struct进行修改则要使用struct的指针最为receiver。

> slice作为receiver
> slice是引用类型，比较特殊


method继承
--------------------

前面了解到字段的继承，其实method也是可以继承的（和匿名字段中的字段一起被引入）。如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该method。
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
    Human  //匿名字段
    school string
}
type Employee struct {
    Human
    company string
}

//在Human上定义一个method，这个method和Human中的字段一样属于Human
func (h *Human) SayHi() {
    fmt.Printf("Hi, I am %s, yon can call me on %s\n", h.name, h.phone)
}
func main() {
    mark := Student{Human{"Mark", 20, "13512344567"}, "MIT"}
    sam := Employee{Human{"Sam", 35, "1701234567"}, "Google Inc"}
    mark.SayHi() //Student和Employee可以直接使用SayHi方法
    sam.SayHi()
}
/*
输出结果：
    Hi, I am Mark, yon can call me on 13512344567
    Hi, I am Sam, yon can call me on 1701234567
*/
```


method重写
-----------------------
上面的例子中，Employee可以实现自己的SayHi方法。和匿名字段冲突一样的道理，我们可以在Employee上面定义一个method，重写匿名字段的方法。如下例，
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
    Human  //匿名字段
    school string
}
type Employee struct {
    Human
    company string
}

//在Human上定义一个method
func (h *Human) SayHi() {
    fmt.Printf("Hi, I am %s, yon can call me on %s\n", h.name, h.phone)
}

//Employee的method重写Human的method
func (e *Employee) SayHi() {
    fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

func main() {
    mark := Student{Human{"Mark", 20, "13512344567"}, "MIT"}
    sam := Employee{Human{"Sam", 35, "1701234567"}, "Google Inc"}

    mark.SayHi() //

    sam.SayHi()    //Employee上的method和匿名字段Human上的method冲突，这样直接调用会首先选择Employee上的method
    sam.Human.SayHi()    //也可以通过这种方式访问匿名字段Human上的method。这和匿名字段冲突一样的道理。
}
/*
输出结果：
    Hi, I am Mark, yon can call me on 13512344567
    Hi, I am Sam, I work at Google Inc. Call me on 1701234567
    Hi, I am Sam, yon can call me on 1701234567
*/
```


**课堂作业**
    根据为结构增加方法的知识，尝试声明一个底层类型为int的类型，并实现调用某个方法就递增100。
```go
package main
import (
    "fmt"
)

type TZ int

func (t *TZ) Increase() {
    *t += 100
}
func main() {
    var a TZ = 110
    a.Increase()
    fmt.Println(a)
}
```
