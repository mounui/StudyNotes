# Redis数据类型

所有的 key 都为`String`类型，讨论数据类型是说的value的类型。

Redis常用5种数据类型：

- string		字符串
- hash 		 哈希
- list              列表
- set              集合
- sorted_set 有序集合

## redis 数据存储格式

- redis 自身是一个 Map，其中所有的数据都是采用 key : value 的形式存储
- 数据类型指的是存储的数据的类型，也就是 value 部分的类型，key 部分永远都是字符串

![2021-06-19_200430](https://gitee.com/mounui/PicBed/raw/master/notebook/2021-06-19_200430.png)

## string 类型

- 存储的数据：单个数据，最简单的数据存储类型，也是最常用的数据存储类型
- 存储数据的格式：一个存储空间保存一个数据
- 存储内容：通常使用字符串，如果字符串以整数的形式展示，可以作为数字操作使用

#### 基本操作

```text
//设置String
set key value
mset key1 value1 key2 value2...
//设置生命周期
setex key seconds value 

//得到String
get key 
mget key1 key2...

//删除String
del key

//向字符串的后面追加字符，如果有就补在后面，如果没有就新建
append key value

//增长指令，只有当value为数字时才能增长
incr key  
incrby key increment  
incrbyfloat key increment 

//减少指令，有当value为数字时才能减少
decr key  
decrby key increment
```

- string在redis内部存储默认就是一个**字符串**，当遇到增减类操作incr，decr时会**转成数值型**进行计算。
- redis所有的操作都是**原子性**的，采用**单线程**处理所有业务，命令是一个一个执行的，因此无需考虑并发带来的数据影响。
- 注意：按数值进行操作的数据，如果原始数据不能转成数值，或超越了redis 数值上限范围，将报错。 9223372036854775807（java中long型数据最大值，Long.MAX_VALUE）

**tips：**

- redis用于控制数据库表主键id，为数据库表主键**提供生成策略**，保障数据库表的主键**唯一性**
- 此方案适用于所有数据库，且支持数据库集群

#### 扩展操作

设置数据具有指定的生命周期

```text
setex key seconds value
psetex key milliseconds value
```

**tips**：redis 控制数据的生命周期，通过数据是否失效控制业务行为，适用于所有具有时效性限定控制的操作

#### 热点数据命名

![2021-06-19_201310](https://gitee.com/mounui/PicBed/raw/master/notebook/2021-06-19_201310.png)

#### 应用场景

1. 具有时效性限定控制的操作
2. 应用于各种结构型和非结构型高热度数据访问加速

## hash 类型

- 新的存储需求：对一系列存储的数据进行编组，方便管理，典型应用存储对象信息
- 需要的存储结构：一个存储空间保存多个键值对数据
- hash类型：底层使用哈希表结构实现数据存储

![2021-06-19_203542](https://gitee.com/mounui/PicBed/raw/master/notebook/2021-06-19_203542.png)

#### 基本操作

```text
//插入（如果已存在同名的field，会被覆盖）
hset key field value
hmset key field1 value1 field2 value2...
//插入（如果已存在同名的field，不会被覆盖）
hsetnx key field value

//取出
hget key field
hgetall key

//删除
hdel key field1 field2...

//获取field数量
hlen key

//查看是否存在
hexists key field

//获取哈希表中所有的字段名或字段值 
hkeys key
hvals key

//设置指定字段的数值数据增加指定范围的值 
hincrby key field increment 
hdecrby key field increment
```

#### 注意事项

- hash类型下的value**只能存储字符串**，不允许存储其他数据类型，**不存在嵌套现象**。如果数据未获取到， 对应的值为（nil）
- 每个 hash 可以存储 2^32 - 1 个键值
- hash类型十分贴近对象的数据存储形式，并且可以灵活添加删除对象属性。但hash设计初衷不是为了存储大量对象而设计的，**切记不可滥用**，更**不可以将hash作为对象列表使用**
- hgetall 操作可以获取全部属性，如果内部field过多，遍历整体**数据效率就很会低**，有可能成为数据访问瓶颈

#### 应用场景

1. 电商网站购物车数据存储设计
2. 抢购，限购类、限量发放优惠卷、激活码等业务的数据存储设计

## list 类型

- 数据存储需求：存储多个数据，并对数据进入存储空间的顺序进行区分
- 需要的存储结构：一个存储空间保存多个数据，且通过数据可以体现进入顺序
- list类型：保存多个数据，底层使用双向链表存储结构实现
- **元素有序，且可重**

![2021-06-19_204331](https://gitee.com/mounui/PicBed/raw/master/notebook/2021-06-19_204331.png)

#### 基本操作

```text
//添加修改数据,lpush为从左边添加，rpush为从右边添加
lpush key value1 value2 value3...
rpush key value1 value2 value3...

//查看数据, 从左边开始向右查看. 如果不知道list有多少个元素，end的值可以为-1,代表倒数第一个元素
//lpush先进的元素放在最后,rpush先进的元素放在最前面
lrange key start end
//得到长度
llen key
//取出对应索引的元素
lindex key index

//获取并移除元素（从list左边或者右边移除）
lpop key
rpop key
```

#### 扩展操作

```text
//规定时间内获取并移除数据,b=block,给定一个时间，如果在指定时间内放入了元素，就移除
blpop key1 key2... timeout
brpop key1 key2... timeout

//移除指定元素 count:移除的个数 value:移除的值。 移除多个相同元素时，从左边开始移除
lrem key count value
```

#### 注意事项

- list中保存的数据都是string类型的，数据总容量是有限的，最多2^32 - 1 个元素 (4294967295)。
- list具有索引的概念，但是操作数据时通常以**队列**的形式进行入队出队(rpush, rpop)操作，或以**栈**的形式进行入栈出栈(lpush, lpop)操作
- 获取全部数据操作结束索引设置为-1 (倒数第一个元素)
- list可以对数据进行分页操作，通常第一页的信息来自于list，第2页及更多的信息通过数据库的形式加载

#### 应用场景

1. 于具有操作先后顺序的数据控制
2. 最新消息展示

## set 类型

- 新的存储需求：存储大量的数据，在查询方面提供更高的效率
- 需要的存储结构：能够保存大量的数据，高效的内部存储机制，便于查询
- set类型：与hash存储结构完全相同，仅存储键，不存储值（nil），并且值是不允许重复的
- **不重复且无序**

![2021-06-19_205049](https://gitee.com/mounui/PicBed/raw/master/notebook/2021-06-19_205049.png)

#### 基本操作

```text
//添加元素
sadd key member1 member2...

//查看元素
smembers key

//移除元素
srem key member

//查看元素个数
scard key

//查看某个元素是否存在
sismember key memberCopy
```

#### 扩展操作

```text
//从set中任意选出count个元素
srandmember key count

//从set中任意选出count个元素并移除
spop key count

//求两个集合的交集、并集、差集
sinter key1 key2...
sunion key1 key2...
sdiff key1 key2...

//求两个set的交集、并集、差集，并放入另一个set中
sinterstore destination key1 key2...
sunionstore destination key1 key2...
sdiffstore destination key1 key2...

//求指定元素从原集合放入目标集合中
smove source destination key
```

#### 注意事项

- set 类型数据操作的注意事项
- set 虽然与hash的存储结构相同，但是无法启用hash中存储值的空间

#### 应用场景

1. 于随机推荐类信息检索，例如热点歌单推荐，热点新闻推荐，热卖旅游线路，应用APP推荐，
   大V推荐等
2. 同类信息的关联搜索，二度关联搜索，深度关联搜索
3. 同类型数据的快速去重
4. 基于黑名单与白名单设定的服务控制

## sorted_set 类型

- 新的存储需求：数据排序有利于数据的有效展示，需要提供一种可以根据自身特征进行排序的方式
- 需要的存储结构：新的存储模型，可以保存可排序的数据
- sorted_set类型：在set的存储结构基础上添加可排序字段
- **不重但有序（score）**

![2021-06-19_205750](https://gitee.com/mounui/PicBed/raw/master/notebook/2021-06-19_205750.png)

#### 基本操作

```text
//插入元素, 需要指定score(用于排序)
zadd key score1 member1 score2 member2

//查看元素(score升序), 当末尾添加withscore时，会将元素的score一起打印出来
zrange key start end (withscore)
//查看元素(score降序), 当末尾添加withscore时，会将元素的score一起打印出来
zrevrange key start end (withscore)

//移除元素
zrem key member1 member2...

//按条件获取数据, 其中offset为索引开始位置，count为获取的数目
zrangebyscore key min max [withscore] [limit offset count]
zrevrangebyscore key max min [withscore] [limit offset count]

//按条件移除元素
zremrangebyrank key start end
zremrangebysocre key min max
//按照从大到小的顺序移除count个值
zpopmax key [count]
//按照从小到大的顺序移除count个值
zpopmin key [count]

//获得元素个数
zcard key

//获得元素在范围内的个数
zcount min max

//求交集、并集并放入destination中, 其中numkey1为要去交集或并集集合的数目
zinterstore destination numkeys key1 key2...
zunionstore destination numkeys key1 key2...
```

#### 注意事项

- score保存的数据存储空间是64位，如果是整数范围是-9007199254740992~9007199254740992
- score保存的数据也可以是一个双精度的double值，基于双精度浮点数的特征，可能会丢失精度，使用时
  候要慎重
- sorted_set 底层存储还是基于set结构的，因此数据不能重复，如果重复添加相同的数据，score值将被反复覆盖，保留最后一次修改的结果

#### 应用场景

1. 计数器组合排序功能对应的排名
2. 定时任务执行顺序管理或任务过期管理
3. 即时任务/消息队列执行管理

## 应用场景总结

1. 控制数据库表主键id，为数据库表主键提供生成策略，保障数据库表的主键唯一性

2. 控制数据的生命周期，通过数据是否失效控制业务行为，适用于所有具有时效性限定控制的操作

3. 各种结构型和非结构型高热度数据访问加速

4. 购物车数据存储设计

5. 抢购，限购类、限量发放优惠卷、激活码等业务的数据存储设计

6. 具有操作先后顺序的数据控制

7. 最新消息展示

8. 随机推荐类信息检索，例如热点歌单推荐，热点新闻推荐，热卖旅游线路，应用APP推荐，大V推荐等

9. 同类信息的关联搜索，二度关联搜索，深度关联搜索

10. 同类型不重复数据的合并、取交集操作

11. 同类型数据的快速去重

12. 基于黑名单与白名单设定的服务控制

13. 计数器组合排序功能对应的排名

14. 定时任务执行顺序管理或任务过期管理

15. 及时任务/消息队列执行管理

16. 按次结算的服务控制

17. 基于时间顺序的数据操作，而不关注具体时间