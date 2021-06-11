# Typora+PicGo+Gitee实现图片上传功能

应用示范

![2021-06-10_233843](https://gitee.com/mounui/PicBed/raw/master/notebook/2021-06-10_233843.gif)

如上，截图粘贴或者直接拖到Typora中即可自动上传并更新图片地址。

## 创建 Gitee 仓库

1. 注册并登录[gitee](https://gitee.com/)
2. 新建仓库用作图床

![image-20210610234031120](https://gitee.com/mounui/PicBed/raw/master/notebook/image-20210610234031120.png)

填写仓库名称和路径

![image-20210610234537812](https://gitee.com/mounui/PicBed/raw/master/notebook/image-20210610234537812.png)

## 生成私人令牌

gitee里点击设置 - 私人令牌 - 生成新令牌

![image-20210610235235025](https://gitee.com/mounui/PicBed/raw/master/notebook/image-20210610235235025.png)

## PicGo 配置 Gitee 图床

1. 下载[PicGo](https://molunerfinn.com/PicGo/)并安装，这是一款图片上传管理工具。
2. 打开PicGo安装Gitee插件（需要使用npm，若没有安装[nodejs](https://nodejs.org/zh-cn/)则安装不了插件）

![Snipaste_2021-06-10_23-09-46](https://gitee.com/mounui/PicBed/raw/master/notebook/Snipaste_2021-06-10_23-09-46.png)

3. 在图床设置中配置Gitee图床

![Snipaste_2021-06-10_23-16-03](https://gitee.com/mounui/PicBed/raw/master/notebook/Snipaste_2021-06-10_23-16-03.png)

点击确定并设置为默认图床

## Typora 配置 Gitee图床

1. 下载安装 [Typora](https://www.typora.io/)
2. 打开 Typora，点击菜单：文件 - 偏好设置 - 图像 配置图床

![2021-06-10_232331](https://gitee.com/mounui/PicBed/raw/master/notebook/2021-06-10_232331.jpg)

!> 点击验证图片上传选项时，会提示使用下面的地址。这里地址端口要保证和 PicGo 中使用的端口一致，如果不一样可以将 PicGo 中的端口修改成这里使用的端口。

![2021-06-10_233843](https://gitee.com/mounui/PicBed/raw/master/notebook/2021-06-10_233843.png)

查看 PicGo 中使用的地址：

点击 PicGo 设置 - 设置 Server，若和 Typora 中图片上传地址不一样，改成 Typora 中的监听端口即可，然后确认。