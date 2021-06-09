# Go语言变量

变量来源于数学，是计算机语言中能储存计算结果或能表示值抽象概念。

变量可以通过变量名访问。

Go 语言变量名由字母、数字、下划线组成，**其中首个字符不能为数字**。

声明变量的一般形式是使用 var 关键字：

```go
var identifier type
// 一次声明多个变量
var identifier1, identifier2 type
```

## 变量声明

#### 1. 指定变量类型

如果没有初始化，则变量默认为零值

```go
var v_name v_type
v_name = value
```

以下类型的零值分别是

- 数值类型（包括complex64/128）为 **0**
- 布尔类型为 **false**
- 字符串为 **""**（空字符串）
- 以下几种类型为 **nil**：

```go
var a *int		// 指针
var a []int		// 切片
var a map[string] int	// 集合
var a chan int		// 通道
var a func(string) int	// 函数
var a error // error 是接口
```

#### 2. 根据值自行判断变量类型

```go
var v_name = value
```

#### 3. 省略var，使用 :=

```go
v_name := value
```

!> := 左侧如果没有声明新的变量，就产生编译错误

```go
var intVal int 
intVal :=1 // 这时候会产生编译错误，因为 intVal 已经声明，不需要重新声明

intVal := 1 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
```

## 多变量声明

```go
package main

var x, y int
var (  // 这种因式分解关键字的写法一般用于声明全局变量
    a int
    b bool
)

//类型相同多个变量, 非全局变量
var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main(){
    g, h := 123, "hello"
    println(x, y, a, b, c, d, e, f, g, h)
}
```

## 值类型和引用类型

值类型：变量直接指向内存中的值，存储在栈中，赋值等同拷贝

引用类型：变量指向内存中的地址，存储在堆中，赋值时会复制引用地址

内存地址（指针）：使用`&`获取变量内存地址，如：`&i`。存储为一个值

## 注意事项

- 不可以重复声明变量
- 不能使用未声明的变量
- 声明的局部变量必须被使用，全局变量可以声明不使用
- 可以简单地使用 **a, b = b, a**交换两个类型相同的变量

!> 标识符 _ 被用于抛弃值，它实际上是个只写变量