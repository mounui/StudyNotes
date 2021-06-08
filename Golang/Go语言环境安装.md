# Go语言环境安装

目前Go语言支持在Linux、Mac OS X、Windows和FreeBSD上安装。开发团队正在尝试将 Go 语言移植到其它例如 OpenBSD、DragonFlyBSD、NetBSD、Plan 9、Haiku 和 Solaris 操作系统上，有关信息可查看：[Go Porting Efforts](http://go-lang.cat-v.org/os-ports)。

## 下载Go

打开[Go语言官网](https://golang.org/dl/)下载，如果因为被墙可以到[镜像网站](https://golang.google.cn/dl/)下载。打开下载页面选择适合自己系统的安装文件下载（这里以go1.16.5为例）各系统对应包名如下

| 操作系统 | 包名                          |
| :------: | :---------------------------- |
| Windows  | go1.16.5.windows-386.msi      |
|  Linux   | go1.16.5.linux-arm64.tar.gz   |
|   Mac    | go1.16.5.darwin-arm64.pkg     |
| FreeBSD  | go1.16.5.freebsd-amd64.tar.gz |

## Windows下安装Go

1. 下载go安装文件
2. 双击安装，如果是默认安装路径一般无需配置环境变量，更改路径需配置好环境变量方便使用
3. 配置GOPATH（工作目录）用户变量指向自己的工作目录，我的电脑->属性->高级系统设置->环境变量添加（D:\goitem 这里是我得工作目录）
4. 打开命令窗口执行`go version`检查是否安装配置成功

## Goland安装

1. 打开[Jetbrain官网](https://www.jetbrains.com/)下载Goland
2. 执行安装

