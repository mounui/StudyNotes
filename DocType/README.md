# docsify常见问题

> docsify安装使用部署参见[docsify文档](https://docsify.js.org/#/zh-cn/)

## docsify-cli安装问题

1. 提示node-gyp，msvs_version等相关问题，命令行错误信息如下：

```
npm ERR! gyp ERR! find VS msvs_version was set from command line or npm config
npm ERR! gyp ERR! find VS - looking for Visual Studio version 2017
npm ERR! gyp ERR! find VS VCINSTALLDIR not set, not running in VS Command Prompt
npm ERR! gyp ERR! find VS could not use PowerShell to find Visual Studio 2017 or newer, try re-running with '--loglevel silly' for more details
npm ERR! gyp ERR! find VS looking for Visual Studio 2015
npm ERR! gyp ERR! find VS - not found
npm ERR! gyp ERR! find VS not looking for VS2013 as it is only supported up to Node.js 8
npm ERR! gyp ERR! find VS
npm ERR! gyp ERR! find VS valid versions for msvs_version:
npm ERR! gyp ERR! find VS
npm ERR! gyp ERR! find VS **************************************************************
npm ERR! gyp ERR! find VS You need to install the latest version of Visual Studio
npm ERR! gyp ERR! find VS including the "Desktop development with C++" workload.
npm ERR! gyp ERR! find VS For more information consult the documentation at:
npm ERR! gyp ERR! find VS https://github.com/nodejs/node-gyp#on-windows
npm ERR! gyp ERR! find VS **************************************************************
```

根据上面错误提示提到的关于Visual Studio地址去查看具体[文档](https://github.com/nodejs/node-gyp#on-windows)，然后根据自己的系统去解决。

这里我使用的是windows，因此我安装了`Visual Studio Build Tools`，然后从新执行`docsify-cli`安装命令就好了。

!> 使用`Visual Studio Installer`安装`Visual Studio 2019`然后打开CMD执行`npm config set msvs_version 2019`