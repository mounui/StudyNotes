# Redis入门

Redis是一种非关系型数据库。

## Redis简介

#### NoSQL

NoSQL：即 Not-Only SQL（ 泛指非关系型的数据库），作为关系型数据库的补充。

作用：应对基于海量用户和海量数据前提下的数据处理问题。

特征：

- 可扩容，可伸缩
- 大数据量下高性能
- 灵活的数据模型
- 高可用

常见 Nosql 数据库：

- **Redis**
- memcache
- HBase
- MongoDB

#### ![2021-06-16_230215](https://gitee.com/mounui/PicBed/raw/master/notebook/2021-06-16_230215.png)Redis

概念：Redis (**RE**mote **DI**ctionary **S**erver) 是用 C 语言开发的一个开源的高性能键值对（**key-value**）数据库。

特征：
1. 数据间没有必然的关联关系
2. 内部采用单线程机制进行工作
3. 高性能。官方提供测试数据，50个并发执行100000 个请求,读的速度是110000 次/s,写的速度是81000次/s。
4. 多数据类型支持
- 字符串类型		string
- 列表类型			list
- 散列类型			hash
- 集合类型			set
- 有序集合类型		sorted_set

5. 持久化支持。可以进行数据灾难恢复

#### Redis的应用

- 为热点数据加速查询（主要场景），如热点商品、热点新闻、热点资讯、推广类等高访问量信息等

- 任务队列，如秒杀、抢购、购票排队等

- 即时信息查询，如各位排行榜、各类网站访问统计、公交到站信息、在线人数信息（聊天室、网站）、设时效性信息控制，如验证码控制、投票控制等

- 分布式数据共享，如分布式集群架构中的 session 分离

- 消息队列

- 分布式锁

## Redis的下载安装

[Redis官网](https://redis.io/)可以下载最新版本，[Windows版本](https://github.com/microsoftarchive/redis/tags)可以在GItHub下载。

#### Windows下安装Redis

![2021-06-18_082926](https://gitee.com/mounui/PicBed/raw/master/notebook/2021-06-18_082926.png)

## Redis基本操作

信息添加：设置 key，value 数据

```redis
set key value	// 命令
set name redis	// 示例
```

信息查询：根据 key 查询对应的 value，如果不存在，返回空（nil）

```redis
get key		// 命令
get name	// 示例
```

清除屏幕信息 `clear`

退出客户端 `quit` `exit` `<ESC>`

帮助：获取命令帮助文档，获取组中所有命令信息名称

```redis
help // 查看帮助信息
help 命令名称
help @组名
```

![2021-06-18_083819](https://gitee.com/mounui/PicBed/raw/master/notebook/2021-06-18_083819.png)

