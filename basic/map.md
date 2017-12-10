Go语言中的map
==================================

	* 类似其他语言中的哈希表或者字典，以key-value形式存储数据
	* key必须是支持==或!=比较运算的类型，不可以是函数、map或slice类型
	* map查找比线性搜索块很多，但比使用索引访问数据的类型慢100倍
	* map使用make()创建，支持“:=”这种简写方式



	* make([keyType] valueType, cap)，cap表示容量，可以省略
	* 超出容量时会自动扩容，但尽量提供一个合理的初始值
	* 使用len()获取元素个数



	* 键值对不存在时自动添加，使用delete()删除某键值对
	* 使用for range对map和slice进行迭代操作


map的几种初始化方式
---------------------------------

```go
func main() {
    var m1 map[int]string
    //m1 = map[int]string{} //或者如下使用make
    m1 = make(map[int]string)
    fmt.Println(m1)

    //var m2 map[int]string = map[int]string{} //或者如下使用make
    var m2 map[int]string = make(map[int]string)
    fmt.Println(m2)

    //更简洁的方式
    m3 := make(map[int]string)
    fmt.Println(m3)
}
```

map的简单操作
----------------------------------
```go
func main() {
    m := make(map[int]string)
    m[1] = "ok"           //插入key=1，value="ok"的数据
    //a := m[1]
    fmt.Println(m[1])  //输出：ok

    delete(m, 1)          //删除
    fmt.Println(m[1])  //没有输出
}
```

复杂map
----------------------------------------

**嵌套的map**

```go
package main
import (
    "fmt"
)
func main() {
    var m map[int]map[int]string
    m = make(map[int]map[int]string)  //这个make只是将最外层的map进行了初始化

    //m[1][1] = "nice"                                //此时直接赋值会报错：panic: runtime error: assignment to entry in nil map，需要对内部的map也进行初始化

    m[1] = make(map[int]string)             //对内部的map进行初始化，注意这里只是对m[1]位置上的map进行初始化，如果要使用m[2]则需要对m[2]位置上的map进行另外初始化。

    //可以使用如下(并行赋值)形式。如果m[2][1]不存在，则a没有值且check被赋值false，否则被赋值true。
    a, check := m[2][1]

    //a, check := m[2][1], "hello"

    //检查check是否为false，可以判断m[2]是否已经被初始化。可以避免出错造成程序停止。
    if !check {
        m[2] = make(map[int]string)
    }

    //m[2][1] = "Good"
    a, check = m[2][1]

    fmt.Println(a)
    fmt.Println(check)
}
```

迭代操作for range
------------------------------------

**slice的迭代**
```go
func main() {
    s := []int{4, 6, 2, 1, 3, 5}
    fmt.Println(s)
    for i, v := range s { //这里的v和i都是局部的，只在for语句块内部有效
        fmt.Print(i)         //i是索引
        fmt.Print(": ")
        fmt.Println(v)     //v是value
    }
}

func main() {
    sm := make([]map[int]string, 5) //创建一个以map为元素的slice
    for _, v := range sm {
        v = make(map[int]string, 1)
        v[1] = "nihao" //这里的v和i都是局部的，只在for语句块内部有效，因此赋值给v并不能改变sm
        fmt.Println(v)
    }
    fmt.Println(sm)  //for中对v进行而来赋值，但原slice中的map仍是空的
}
/*
输出结果：
map[1:nihao]
map[1:nihao]
map[1:nihao]
map[1:nihao]
map[1:nihao]
[map[] map[] map[] map[] map[]]
/*
```

**map的迭代**
```go
package main
import (
    "fmt"
)
func main() {     
    m := map[int]string{1: "hi", 2: "hello", 3: "nihao", 4: "bonjor", 5: "an"}
    for k, v := range m {     //map的迭代是无序的
        fmt.Print(k)                //k是键
        fmt.Print(": ")
        fmt.Println(v)             //v是值
    }
}
/*
    3: nihao
    4: bonjor
    5: an
    1: hi
    2: hello
*/
```

**顺序输出一个map**
    把map的key存储到slice中，对slice进行排序，然后按slice中k的顺序取出map中对应的值。
```go
package main
import (
    "fmt"
    "sort"
)
func main() {     //对一个map的间接排序
    m := map[int]string{1: "hi", 2: "hello", 3: "nihao", 4: "bonjor", 5: "an"}
    s := make([]int, len(m))
    i := 0
    for k, _ := range m {
        s[i] = k
        i++
    }
    sort.Ints(s)
    fmt.Println(s)
}
```
课堂作业
--------------------------------
根据在for range的知识，尝试将类型为map[int] string的键和值进行交换，变成类型map[string]int.
```go
package main

import (
    "fmt"
)

func main() {
    m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "3"}
    fmt.Println(m1)

    m2 := make(map[string]int)

    for k, v := range m1 {
        m2[v] = k
    }
    fmt.Println(m2)
}
```
