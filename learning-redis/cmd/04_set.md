# set

# SADD

格式: `SADD key member1 [member2 member3 ...]`

说明: 将一个或多个 member 加入到集合 key 中, 已存在的 member 忽略.

返回: 被添加到集合中的 member 数量


# SISMEMBER

格式: `SISMEMBER key member`

说明: 判断 member 是否为 key 的成员

返回: 是返回 `1`, 否返回 `0`


# SPOP

格式: `SPOP key`

说明: 随机移除集合 key 中的一个元素

返回: 移除的元素, 集合不存在或为空时返回 `nil`


# SRANDMEMBER

格式: `SRANDMEMBER key [count]`

说明: 从集合 key 中随机返回元素; 不提供 count 时返回单个元素;
count > 0 时返回 count 个集合 key 中的元素. 若 count > len(key) 返回整个集合
count < 0 时返回 abs(count) 个集合 key 中的元素, 可能有重复.

返回: 无 count 返回单个元素, 集合为空返回 `nil`, 如果提供 count 则返回列表


# SREM

格式: `SREM key member1 [member2 member3 ...]`

说明: 移除 key 中的一到多个元素, 不存在的会被忽略

返回: 成功移除的元素数量


# SMOVE

格式: `SMOVE source destination member`

说明: 将 member 从 source 移动到 destination

返回: source 不存在或不包含 member 则不执行任何操作, 返回 `0`, 成功返回 `1`


# SCARD

格式: `SCARD key`

说明: 返回集合 key 的元素数量

返回: key 的元素数量, 不存在为 0


# SMEMBERS

格式: `SMEMBERS key`

说明: 返回集合 key 中的所有成员

返回: 集合中的所有成员


# SSCAN

格式: `SSCAN key cursor [MATCH patten][COUNT count]`

说明: 扫描集合返回符合条件的元素

返回: 二元组, 第一个为游标, 为 0 表示结束迭代, 其他值表示未结束. 第二个是元素列表.


# SINTER

格式: `SINTER key1 [key2 key3 ...]`

说明: 返回给定所有集合的交集

返回: 交集的成员列表


# SINTERSTORE

格式: `SINTERSTORE destination key1 [key2 key3 ...]`

说明: 类似 SINTER, 将结果存入 destination, 覆盖原来的值

返回: 返回交集成员数量


# SUNION

格式: `SUNION key1 [key2 key3 ...]`

说明: 返回给定所有集合的并集

返回: 并集的成员列表


# SUNIONSTORE

格式: `SUNIONSTORE destination key1 [key2 key3 ...]`

说明: 类似 SUNION, 将结果存入 destination, 覆盖原来的值

返回: 返回并集成员数量


# SDIFF

格式: `SDIFF key1 [key2 key3 ...]`

说明: 返回给定所有集合的差集

返回: 差集的成员列表


# SDIFFSTORE

格式: `SDIFFSTORE destination key1 [key2 key3 ...]`

说明: 类似 SDIFF, 将结果存入 destination, 覆盖原来的值

返回: 返回差集成员数量
