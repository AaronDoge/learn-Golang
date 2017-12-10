Go语言切片Slice
=====================================

	* Slice本身并不是数组，它指向底层的数组
	* 作为变长数组的替代方案，可以关联底层数组的局部或全部
	* 为引用类型
	* 可以直接创建或从底层数组获取生成
	* 使用len()获取元素个数，cap()获取容量
	* 一般使用make()创建
	* 如果多个slice指向同一个底层数组，其中一个的值改变会影响全部


创建切片
------------------------------------

	* make([]T, len, cap)
	* 其中cap可以省略，则默认和len相等
	* len表示元素个数，cap表示容量

	* 也可以通过 var s1 []int 的方式来定义一个slice


 **数组中截取部分数据：**
```go
func main() {
    //var s1 []int
    a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    fmt.Println(a)

    s1 := a[0:5]        //获取第0位到第5位(不包括第5个)元素，左闭右开。
    s3 := a[:5]          //同上
    fmt.Println(s1)  //[0 1 2 3 4]
    fmt.Println(s3)  //[0 1 2 3 4]

    s2 := a[5:]          //获取第5位到最后一个元素
    s2 := a[5:len(a)] //同上
    fmt.Println(s2)  //[5 6 7 8 9]
}
```

Reslice
----------------------------------
	* Reslice时索引以被slice的切片为准
	* 切片的两个参数len和cap，切片数据的大小，cap表示切片起始位置到被切数组末尾的元素个数
	* 索引不可以超过被slice的切片的容量cap()值
	* Reslice时，当slice的位置超过len时（当然必须是小于cap的），则取被切数组对应位的数据
	* 索引越界不会导致底层数组的重新分配而是引发错误

```go
package main
import (
    "fmt"
)
func main() {
    //定义slice
    s1 := make([]int, 3, 10)
    fmt.Println(s1)

    a := [...]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l'}

    sa := a[2:5]
    fmt.Println(string(sa))               //cde
    fmt.Println(len(sa), cap(sa))     //len=3, cap=10

    //reslice
    sb := sa[1:3]
    fmt.Println(string(sb))     //输出de
    sc := sa[1:5]
    fmt.Println(string(sc))     //输出defg，这个地方注意，reslice位置大于slice的len，则取被切数组对应位的数据

    sb[1] = 'x'             //这里的修改相当于修改数组a的数据
    fmt.Println(string(sb))     //dx
    fmt.Println(string(sc))     //dxfg
    //fmt.Println(string(a))     //为啥string(a)会报错？？这里也是数组和slice的区别之一
    fmt.Println(a)     //[97 98 99 100 120 102 103 104 105 106 107 108],可以看到对应位置已经被修改。
}
```


切片的append()函数
--------------------------------------------

格式：`slice2 = append(slice1, data1, data2, ...)`，向slice1中一次append数据data1，data2，...，并返回slice2。
```go
package main
import (
    "fmt"
)
func main() {
    a := []int{1, 2, 3, 4, 5}
    s1 := a[2:5]
    s2 := a[1:3]
    fmt.Println(s1, s2)     //[3 4 5] [2 3]

    s1[0] = 7
    fmt.Println(s1, s2)     //[7 4 5] [2 7] 可以发现上一行通过slice_1修改了数组a的值，在s2中也得到体现

    s3 := append(s2, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6)
    s1[0] = 9
    fmt.Println(s1, s2, s3)     //[9 4 5] [2 9] [2 7 6 6 6 6 6 6 6 6 6 6 6] 可以发现s3中的对应位并没有改变，为什么？
    //这是因为在对slice进行append操作时，如果append后slice的cap大小超过原来数组的大小(从slice截取开始处计)时，会新建一个数组，这时此slice不再指向原数组额，而是指向新数组。
}
```
 
copy(s1, s2)函数
-----------------------------------
后者数据拷贝到前者，即s2的数据拷贝到s1中。
```go
func main() {
    s1 := []int{1, 2, 3, 4, 5, 6}
    s2 := []int{7, 8, 9}

    //copy(s1, s2)：s2向s1拷贝数据
    copy(s1, s2)
    fmt.Println(s1)     //输出：[7 8 9 4 5 6]

    s3 := []int{1, 2, 3, 4, 5, 6}
    s4 := []int{7, 8, 9}
    copy(s4, s3)
    fmt.Println(s4)     //输出：[1 2 3]
    //后者向前者拷贝数据
    //少的向多的拷贝时->替换
    //多的向少的拷贝时->截取

    s5 := []int{1, 2, 3, 4, 5, 6}
    s6 := []int{7, 8, 9}
    copy(s5[2:4], s6[1:3])
    fmt.Println(s5)     //输出：[1 2 8 9 5 6] 替换指定位置
}
```

课堂作业
-------------------------------
如何将一个slice指向一个完整的底层数组，而不是底层数组的一部分。
```go
package main
import (
    "fmt"
)
func main() {
    s1 := []int{1, 2, 3, 4, 5, 6}
    //s2 := []int{7, 8, 9}
    //num := len(s1)
    fmt.Println(s1)

    //采用copy的方式
    var s3 = make([]int, len(s1))
    copy(s3, s1)
    fmt.Println(s3)

    //直接赋值
    s4 := s1
    fmt.Println(s4)
    //fmt.Println(s1 == s4)  //报错： slice can only be compared to nil   
    fmt.Println(s1 == nil)     //false

    //采用slice的方式，从头到尾进行切片的情况可以省略slice的开头和结尾
    s5 := s1[:]
    fmt.Println(s5)
}
```
