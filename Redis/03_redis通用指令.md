# Redis通用指令

## key 通用指令

key特征：key是一个字符串，通过key获取redis中保存的数据

#### key 基本操作

```text
del key		// 删除指定key
exists key	// 获取key是否存在
type key	// 获取key的类型
```

#### key 时效性控制

```text
// 为指定key设置有效期
expire key seconds
pexpire key milliseconds
expireat key timestamp
pexpireat key milliseconds-timestamp
// 获取key的有效时间
ttl key
pttl key
// 切换Key从时效性转换为永久性
persist key
```

#### key 查询模式

```text
// 查询key
keys pattern

// 查询模式规则
* 	匹配任意数量的任意符号
? 	匹配一个符号
[] 	匹配一个指定符号

eg:
keys * 		查询所有
keys it* 	查询所有以it开头
keys *heima 	查询所有以heima结尾
keys ??heima 	查询所有前面两个字符任意，后面以heima结尾
keys user:? 	查询所有以user:开头，最后一个字符任意
keys u[st]er:1 	查询所有以u开头，以er:1结尾，中间包含一个字母，s或t
```

#### key 其他操作

```text
// 重命名key，为了避免覆盖已有数据，尽量少去修改已有key的名字，如果要使用最好使用renamenx
rename key newKey
renamenx key newKey

// 对所有key排序
sort

// 查看所有关于key的操作, 可以使用Tab快速切换
help @generic
```

## 数据库通用指令

#### 数据库

- Redis为每个服务提供有16个数据库，编号从0到15
- 每个数据库之间的数据相互独立

#### db 基本操作

```text
// 切换数据库
select index

// 其他操作
quit
ping
echo message
```

#### db 相关操作

```text
// 移动数据, 必须保证目的数据库中没有该数据
move key db

dbsize	// 查看该库中数据总量
flushdb		// 清除当前数据库
flushall	// 清除所有数据库
```

