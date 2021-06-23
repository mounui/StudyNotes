# Redis持久化

## 简介

什么是持久化：利用永久性存储介质将数据进行保存，在特定的时间将保存的数据进行恢复的工作机制称为持久化。

为什么要持久化：防止数据的意外丢失，确保数据安全性

**持久化过程保存什么：**

- 将当前数据状态进行保存，快照形式，存储数据结果，存储格式简单，关注点在数据
- 将数据的操作过程进行保存，日志形式，存储操作过程，存储格式复杂，关注点在数据的操作过程

![2021-06-23_082735](https://gitee.com/mounui/PicBed/raw/master/notebook/2021-06-23_082735.png)

## RDB

#### RDB启动方式 save

- 命令

  ```
  saveCopy
  ```

- 作用

  手动执行一次保存操作

#### RDB配置相关命令

- dbfilename dump.rdb
  - 说明：设置本地数据库文件名，默认值为 dump.rdb
  - 经验：通常设置为dump-端口号.rdb
- dir
  - 说明：设置存储.rdb文件的路径
  - 经验：通常设置成存储空间较大的目录中，目录名称data
- rdbcompression yes
  - 说明：设置存储至本地数据库时是否压缩数据，默认为 yes，采用 LZF 压缩
  - 经验：通常默认为开启状态，如果设置为no，可以节省 CPU 运行时间，但会使存储的文件变大（巨大）
- rdbchecksum yes
  - 说明：设置是否进行RDB文件格式校验，该校验过程在写文件和读文件过程均进行
  - 经验：通常默认为开启状态，如果设置为no，可以节约读写性过程约10%时间消耗，但是存储一定的数据损坏风险

#### save指令工作原理

![2021-06-23_083655](https://gitee.com/mounui/PicBed/raw/master/notebook/2021-06-23_083655.png)

!> save指令的执行会阻塞当前Redis服务器，直到当前RDB过程完成为止，有可能会造成长时间阻塞，线上环境不建议使用。

#### RDB启动方式 bgsave

- 命令

  ```
  bgsaveCopy
  ```

- 作用

  手动启动后台保存操作，但**不是立即执行**

#### bgsave指令工作原理

![](https://gitee.com/mounui/PicBed/raw/master/notebook/2021-06-23_083839.png)

!> bgsave命令是针对save阻塞问题做的优化。Redis内部所有涉及到RDB操作都采用bgsave的方式，save命令可以放弃使用，推荐使用bgsave

