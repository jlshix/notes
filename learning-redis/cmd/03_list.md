# list

## LPUSH

格式: `LPUSH key value1 [value2 value3 ...]`

说明: 将一到多个 value 插入到列表 key 的表头; 需要注意的是, 插入多个时, 结果与输入顺序相反.
如 `LPUSH mylist a b c` 后 mylist 的内容为 `c b a`

返回: 执行后列表的长度


## LPUSHX

格式: `LPUSHX key value`

说明: 仅当 key 存在且为列表时将 value 插入到 key 的表头

返回: 执行后列表的长度


## RPUSH

格式: `RPUSH key value1 [value2 value3 ...]`

说明: 将一到多个 value 插入到列表 key 的表尾

返回: 执行后列表的长度


## RPUSHX

格式: `RPUSHX key value`

说明: 仅当 key 存在且为列表时将 value 插入到 key 的表尾

返回: 执行后列表的长度


## LPOP

格式: `LPOP key`

说明: 移除并返回 key 的头元素

返回: 列表的头元素, key 不存在时返回 `nil`


## RPOP

格式: `RPOP key`

说明: 移除并返回 key 的尾元素

返回: 列表的尾元素, key 不存在时返回 `nil`


## RPOPLPUSH

格式: `RPOPLPUSH source destination`

说明: 弹出 source 表尾元素插入 destination 的表头. source 和 destination 可为同一个
列表, 此时执行 rotation 操作

返回: source 不存在返回 `nil`, 成功返回此元素


## LREM

格式: `LREM key count value`

说明: 根据参数 count 的值移除列表中与参数 value 相等的元素.
count > 0 时从头到尾移除 count 个与 value 相等的元素;
count < 0 时从为到头移除 abs(count) 个与 value 相等的元素;
count == 0 时移除表中所有与 value 相等的值.

返回: 被移除的元素的数量, key 不存在时也返回 0


## LLEN

格式: `LLEN key`

说明: 返回列表 key 的长度

返回: 返回列表 key 的长度, key 不存在返回 0


## LINDEX

格式: `LINDEX key index`

说明: 返回列表 key 中下标为 index 的元素

返回: 返回列表 key 中下标为 index 的元素, 若 index 不在 key 的索引范围内, 返回 `nil`


## LINSERT

格式: `LINSERT key BEFORE|AFTER pivot value`

说明: 将 value 插入到列表 key 中, 位于 pivot 之前或之后, key 或 pivot 不存在时不执行任何操作

返回: 成功返回列表长度; 无 pivot 返回 `-1`; 无 key 返回 `0`


## LSET

格式: `LSET key index value`

说明: 将列表 key 索引为 index 的值设为 value

返回: 成功返回 `OK`, key 不存在或 index 超出索引返回时返回 `error`


## LRANGE

格式: `LRANGE key start stop`

说明: 返回列表 key 中指定区间内的元素列表

返回: 返回列表 key 中指定区间内的元素列表, 超出索引范围不会引发异常


## LTRIM

格式: `LTRIM key start stop`

说明: 对一个列表进行修剪

返回: 执行成功后返回 OK, 超出索引范围不会引发异常


## BLPOP

格式: `BLPOP key1 [key2 key3 ...] timeout`

说明: 列表阻塞式弹出, 为 `LPOP key` 的阻塞版本. key 为多个时按先后检查, 弹出第一个
非空的头元素. timeout 为 0 时无限期等待. 在 `MULTI/EXEC` 事务中与 `LPOP` 无异.

返回: 列表为空返回一个 `nil`, 否则返回一个二元组, 为列表名和弹出的头元素


## BRPOP

格式: `BRPOP key1 [key2 key3 ...] timeout`

说明: 列表阻塞式弹出, 为 `RPOP key` 的阻塞版本. key 为多个时按先后检查, 弹出第一个
非空的尾元素. timeout 为 0 时无限期等待. 在 `MULTI/EXEC` 事务中与 `RPOP` 无异.

返回: 列表为空或超时返回一个 `nil`, 否则返回一个二元组, 为列表名和弹出的尾元素


## BRPOPLPUSH

格式: `BRPOPLPUSH source destination`

说明: `RPOPLPUSH` 的阻塞版本, source 不为空时二者一致.

返回: 列表为空或超时返回 `nil` 和等待时长, 否则返回被弹出元素的值和等待时长.
