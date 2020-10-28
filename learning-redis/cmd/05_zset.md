# zset

# ZADD

格式: `ZADD key score1 member1 [score2 member2 ...]`

说明: 将一到多个分值和元素放入有序集合 key 中. 若某元素已存在, 更新其 score.
score 可以是整数或双精度浮点数.

返回: 成功添加的元素数量, 不包含已更新的


# ZSCORE

格式: `ZSCORE key member`

说明: 获取 key 中 member 的 score

返回: score 值, 不存在返回 `nil`


# ZINCRBY

格式: `ZINCRBY key increment member`

说明: 为 key 的成员 member 的 score 加上 increment. 不存在时相当于 ZADD

返回: 新的 score 值


# ZCARD

格式: `ZCARD key`

说明: 返回 key 的元素数量

返回: 返回 key 的元素数量, 不存在则返回 `0`


# ZCOUNT

格式: `ZCOUNT key min max`

说明: 返回 key 中 min <= score <= max 元素的数量

返回: key 中 min <= score <= max 元素的数量


# ZRANGE

格式: `ZRANGE key start stop [WITHSCORES]`

说明: 返回 key 中指定区间内的成员, 位置按 score 值递增排序, score 相同时按字典序

返回: 符合条件的成员列表


# ZREVRANGE

格式: `ZREVRANGE key start stop [WITHSCORES]`

说明: 返回 key 中指定区间内的成员, 位置按 score 值递减排序, score 相同时按字典逆序

返回: 符合条件的成员列表


# ZRANGEBYSCORE

格式: `ZRANGEBYSCORE key min max [WITHSCORES][LIMIT offset count]`

说明: 类似 ZRANGE, min max 表示 score 区间, 可选的 LIMIT 参数表示结果的数量和区间.
另外, min 和 max 可以是 `-inf` 和 `+inf`; `(start` `(stop` 可以表示开区间

返回: 符合条件的成员列表


# ZREVRANGEBYSCORE

格式: `ZREVRANGEBYSCORE key max min [WITHSCORES][LIMIT offset count]`

说明: 类似 ZRANGEBYSCORE, 注意参数 max 在前 min 在后, 逆序排列

返回: 符合条件的成员列表


# ZRANK

格式: `ZRANK key member`

说明: 返回 key 中 member 的按 score 从小到大的排名, 最小为 0

返回: key 中 member 的排名, 不在返回 `nil`


# ZREVRANK

格式: `ZREVRANK key member`

说明: 类似 ZRANK, 顺序为从大到小

返回: key 中 member 的排名, 不在返回 `nil`


# ZREM

格式: `ZREM key member1 [member2 member3 ...]`

说明: 移除有序集中一到多个成员, 忽略不存在的

返回: 成功移除的成员数量


# ZREMRANGEBYRANK

格式: `ZREMRANGEBYRANK key start stop`

说明: 移除有序集中指定排名区间内的所有成员

返回: 被移除的成员数量


# ZREMRANGEBYSCORE

格式: `ZREMRANGEBYSCORE key start stop`

说明: 移除有序集中指定分值区间内的所有成员

返回: 被移除的成员数量


# ZRANGEBYLEX

格式: `ZRANGEBYLEX key min max [LIMIT offset count]`

说明: 当有序集he所有成员都具有相同的分值时, 所有元素根据字典序排列. 返回所有值
介于 min 和 max 之间的成员. 合法的 min 和 max 必须包含 `(` 或 `[` 表示开闭区间.
`-` 和 `+` 表示负无限和正无限

返回: 符合条件的成员列表


# ZLEXCOUNT

格式: `ZLEXCOUNT key min max`

说明: 类似 ZRANGEBYLEX, 返回数量

返回: 符合条件的成员数量


# ZREMRANGEBYLEX

格式: `ZREMRANGEBYLEX key min max`

说明: 类似 ZRANGEBYLEX, 移除符合条件的元素

返回: 被移除的元素数量


# ZSCAN

格式: `ZSCAN key cursor [MATCH patten][COUNT count]`

说明: 扫描有序集, 返回符合条件的元素

返回: 二元组, 第一个为游标, 为 0 表示结束迭代, 其他值表示未结束. 第二个是元素列表.


# ZUNIONSTORE

格式: `ZUNIONSTORE destination numkeys key1 [key2 ...] [WEIGHTS weight1 [weight2 ...]] [AGGREGATE SUM|MIN|MAX]`

说明: 计算给定的一到多个有序集的并集, numkeys 为 key 的数量, 结果存到 destination

参数:
- WEIGHTS: 指定每个 key 的乘法因子, key 的 score 传递到聚合函数前先乘此因子, 默认为 1

- AGGREGATE: 并集结果集的聚合方式, SUM 为所有之和, MIN 和 MAX 分别取最小最大值

返回: destination 中结果的数量


# ZINTERSTORE

格式: `ZINTERSTORE destination numkeys key1 [key2 ...] [WEIGHTS weight1 [weight2 ...]] [AGGREGATE SUM|MIN|MAX]`

说明: 类似 ZUNIONSTORE, 计算交集

返回: destination 中结果的数量

