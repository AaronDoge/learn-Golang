Go语言数组
===============================

	* 定义数组的格式：var <varName> [n] <type>，其中n>=0
	* 数组长度也是类型的一部分，因此具有不同长度的数组为不同类型
	* 注意区分指向数组的指针和指针数组
	* 数组在Go中为值类型
	* 数组之间可以使用==或!=进行比较，但不可以使用<或>
	* 可以使用new来创建数组，此方法返回一个指向数组的指针
	* Go支持多维数组


一维数组
-------------------------------------
```go
package main

import (
    "fmt"
)

func main() {
    //var aa [2]int
    //var ab [1]int //数组长度是数组类型的一部分，ax和bx长度不同，是不同类型数组

    x, y := 1, 2

    ac := [5]int{1}     //[1 0 0 0 0]
    ac[2] = 2             //[1 0 2 0 0]
    ad := [20]int{19: 1}              //[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]
    ae := [...]int{1, 2, 3, 4, 5}     //[1 2 3 4 5]
    af := [...]int{19: 1}                //[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]

    var pa *[20]int = &af         //输出pa：&[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]
                                                 //输出*pa：[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]

    //pa为数组指针，pb为指针数组
    pb := [...]*int{&x, &y} //[0xc0420120a8 0xc0420120c0]

    //使用new创建，返回一个指向数组的指针
    pc := new([10]int)
    pc[1] = 2     //即使pc是一个指针，同样也可以使用这种形式给数组赋值

    fmt.Println(ac)
    fmt.Println(ad)
    fmt.Println(ae)
    fmt.Println(af)
    fmt.Println(*pa)         //[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]
    fmt.Println(pb)          //[0xc0420120a8 0xc0420120c0]
    fmt.Println(ad == af)     //true，Go中可以直接进行数组比较
    fmt.Println(pc)         //&[0 2 0 0 0 0 0 0 0 0]
}
```

多维数组
-------------------------------------
```go
package main
import (
    "fmt"
)
func main() {
    a := [2][3]int{
        {1, 1, 1},
        {2, 2, 2}}     //Go中对括号位置敏感，这里尾括号单放一行会报错。

    b := [2][3]int{{1: 1}, {1: 2}}
    c := [...][3]int{{1: 1}, {1: 2}}

    //d := [...][...]int{{1: 1}, {1: 2}}     //这种情况会报错

    fmt.Println(a)     //[[1 1 1] [2 2 2]]
    fmt.Println(b)     //[[0 1 0] [0 2 0]]
    fmt.Println(c)     //[[0 1 0] [0 2 0]]
}
```

Go实现冒泡排序
---------------------------------------
```go
package main
import (
    "fmt"
)
func main() {
    a := [...]int{2, 5, 1, 8, 4}
    fmt.Println("排序前：")
    fmt.Println(a)
    num := len(a)
    for i := 0; i < num; i++ {
        for j := i + 1; j < num; j++ {
            if a[i] < a[j] {
                temp := a[i]
                a[i] = a[j]
                a[j] = temp
            }
        }
    }
    fmt.Println("排序后：")
    fmt.Println(a)
}
```

