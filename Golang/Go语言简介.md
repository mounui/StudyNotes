# Go语言简介

Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。

Go是从2007年末由Robert Griesemer, Rob Pike, Ken Thompson主持开发，后来还加入了Ian Lance Taylor, Russ Cox等人，并最终于2009年11月开源，在2012年早些时候发布了Go 1稳定版本。现在Go的开发已经是完全开放的，并且拥有一个活跃的社区。

## 特点

- 简洁、快速、安全
- 并行、有趣、开源
- 内存管理、数组安全、编译迅速

## 用途

- 搭载 Web 服务器，存储集群或类似用途的巨型中央服务器
- 高性能分布式系统，游戏端服务器开发

## 第一个Go程序

第一个 Go 程序 hello.go（Go 语言源文件的扩展是 .go）

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

执行：使用 **go run** 命令执行程序

```shell
$ go run hello.go 
Hello, World!
```

编译：使用 **go build** 命令来生成二进制文件

```shell
$ go build hello.go 
$ ls
hello    hello.go
$ ./hello 
Hello, World!
```

