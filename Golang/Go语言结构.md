# Go语言结构

Go 语言的基础组成有：包声明、引入包、函数、变量、语句 & 表达式、注释

## 代码示例

```go
// 包声明 必须在源文件中非注释的第一行指明这个文件属于哪个包
package main

import "fmt"	// 引入包 需要使用

// 函数 main函数是每一个可执行程序所必须包含的，启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）
func main() {
	var hello string = "Hello "		// 变量
	world := "World!"	// 变量
	variable := hello + world 	// 表达式
	/*这是一个简单的程序*/
	fmt.Println(variable)
}
```

以一个大写字母开头的标识符可以被外部包使用，这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。

## 执行程序

1. 将以上代码保存为*hello.go*
2. 打开命令行进入文件保存目录
3. 输入命令 *go run hello.go* 并按回车执行代码
4. 如果操作正确你将在屏幕上看到 *"Hello World!"* 字样的输出

```shell
$ go run hello.go
Hello World!
```

5. 还可以使用 **go build** 命令来生成二进制文件

```shell
$ go build hello.go 
$ ls
hello    hello.go
$ ./hello 
Hello, World!
```

!> 左括号 **{** 不能单独放在一行，以下代码在运行时会产生错误：

```go
package main

import "fmt"

func main()  
{  // 错误，{ 不能在单独的行上
    fmt.Println("Hello, World!")
}
```





