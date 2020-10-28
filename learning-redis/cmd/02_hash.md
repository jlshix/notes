# hash

## HSET

格式: `HSET hash field value`

说明: 将 hash 的 field 设为 value, 不存在则创建, 存在则覆盖

返回: 新建 field 并成功设置 value 返回 `1`; field 已存在并成功覆盖旧值返回 `0`


## HSETNX

格式: `HSET hash field value`

说明: 仅当 field 不存在于 hash 时, 设置值 value

返回: 设置成功返回 `1`, 已存在放弃执行返回 `0`


## HGET

格式: `HGET hash field`

说明: 获取 hash 的 field 的值

返回: 有值返回值, 无值返回 `nil`


## HEXISTS

格式: `HEXISTS hash field`

说明: 检查 field 是否存在于 hash 中

返回: 存在返回 `1`, 不存在返回 `0`


## HDEL

格式: `HDEL hash field1 [field2 field3 ...]`

说明: 删除 hash 中的一个或多个指定 field, 不存在的将被忽略

返回: 成功删除的 field 数量


## HLEN

格式: `HLEN hash`

说明: 查询 hash 的 field 数量

返回: hash 的 field 数量, hash 不存在时返回 `0`


## HSTRLEN

格式: `HSTRLEN hash field`

说明: 查询 hash 的 field 对应至的字符串长度

返回: 字符串的长度, 不存在则为 `0`


## HINCRBY

格式: `HINCRBY hash field increment`

说明: 为 hash 的 field 的值加上增量 increment, hash 或 field 不存在时新建

返回: 不为数字时返回 `error`, 否则返回命令执行后的值


## HINCRBYFLOAT

格式: `HINCRBYFLOAT hash field incrment`

说明: HINCRBY 的浮点数版本

返回: 不为数字时返回 `error`, 否则返回命令执行后的值


## HMSET

格式: `HMSET hash field1 value1 [field2 value2 ...]`

说明: 同时将多个键值对设置到 hash, 存在则覆盖, 不存在则创建

返回: 成功返回 `OK`, hash 不为哈希表时返回 `error`


## HMGET

格式: `HMGET hash field1 [field2 field3 ...]`

说明: 返回哈希表 hash 中一个或多个给定 field 的值, 若某 field 不存在则对应为 `nil`

返回: 对应 field 的 value 列表


## HKEYS

格式: `HKEYS hash`

说明: 返回哈希表 hash 中所有的 field

返回: hash 的 field 列表, hash 不存在时返回空列表


## HVALS

格式: `HVALS hash`

说明: 返回哈希表 hash 中所有的 value

返回: hash 的 value 列表, hash 不存在时返回空列表


## HGETALL

格式: `HGETALL hash`

说明: 返回哈希表 hash 中所有的 field value

返回: 以列表形式返回所有的 field 和 value 例 `field1 value1 field2 value2 ...`


## HSCAN

格式: `HSCAN hash cursor [MATCH patten][COUNT count]`

说明: 迭代哈希表的键值对

返回: 二元组, 第一个为游标, 为 0 表示结束迭代, 其他值表示未结束. 第二个是键值对列表.



