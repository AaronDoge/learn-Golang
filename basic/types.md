Go语言基本类型
=================================

基本数据类型包括
---------------------------------

	* bool型：bool

        -长度：1字节
        -取值范围：true，false
        -注意事项：不可以用数组代表true或false

	* 整型：int/uint

        -根据运行平台可能为32位或64位

	* 8位整型：int8/uint8

        -长度：1字节
        -取值范围：-128~127/0~255

	* 字节型：byte（uint8别名）



	* 16位整型：int16/uint16

        -长度：2字节
        -取值范围：-32768~32767/0~65535

	* 32位整型：int32（rune）/uint32

        -长度：8字节
        -取值范围：-2^64/2~(2^64/2-1)/0~(2^64-1)

	* 64位整型：int64/uint64

        -长度：8字节
        -取值范围：-2^64/2~(2^64/2-1)/0~(2^64-1)

	* 浮点型：float32/float64

        -长度：4/8字节
        -小数位：精确到7/15小数位


	* 复数：complex64/complex128

        -长度：8/16字节
        -足够保存指针的32位或64位整数型：uintptr

其他值类型：`array`、`struct`、`string`
引用类型：`slice`、`map`、`chan`

接口类型：`interface`
函数类型：`func`

类型零值
-----------------------------------

零值并不等于空值，而是当变量被声明为某种类型后的默认值，通常情况下值类型的默认值为0，bool为false，string为空字符串。

```go
package main
import (
    "fmt"
    "math"
)
func main() {
    var a [3]byte
    fmt.Println(a)
    fmt.Println(math.MaxInt16)
    fmt.Println(math.MinInt16)
}
```

类型别名
------------------------------

```go   
package main
import (
    "fmt"
    "math"
)
//定义类型别名
type (
    byte       int8
    rune       int32
    文本         string
    BigInteger int64
)
func main() {
    var a [3]byte
    fmt.Println(a)
    fmt.Println(math.MaxInt16)
    fmt.Println(math.MinInt16)

    var b 文本
    b = "中文类型，体现了go支持UTF8的优势"
    fmt.Println(b)

    var c BigInteger
    fmt.Println(c)
}
```
    
从严格意义上讲`type newInt int`，这里的newInt并不能硕士int的别名，而只是底层数据结构相同，在这里称为自定义类型，在进行类型转换时仍需要显式转换，但byte和rune确确实实为uint8和int32的别名，可以相互进行转换。
