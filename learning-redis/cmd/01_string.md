# string 字符串命令

## SET

格式: `SET key value [EX seconds][PX milliseconds][NX|XX]`

说明: 将 key 设为 value.

参数:
- `EX seconds`: 将过期时间设为 seconds 秒, `set key value ex seconds`
等价于 `setex key seconds value`

- `PX seconds`: 将过期时间设为 milliseconds 秒, `set key value px milliseconds`
等价于 `psetex key milliseconds value`

- `NX`: `set key value nx` 相当于 `setnx key value`

- `XX`: key 已存在时才进行 SET 操作

返回: 成功返回 `OK`, 失败返回 `nil`


## SETNX

格式: `SETNX key value`

说明: 只有在 key 不存在的情况下, 才将 key 的值设为 value

返回: 成功返回 `1`, 失败返回 `0`


## SETEX

格式: `SET key seconds value`

说明: 将 key 的值设为 value, 并将 key 的生存时间设为 seconds 秒钟

返回: 成功返回 `OK`, seconds 参数不合法返回 `error`


## PSETEX

格式: `PSETEX key milliseconds value`

说明: SETEX 的毫秒版本

返回: 成功返回 `OK`


## GET

格式: `GET key`

说明: 获取 key 的值

返回: 存在则返回 key 的值, 不存在返回 `nil`, key 不为字符串返回 `error`


## GETSET

格式: `GETSET key value`

说明: 将 key 的值设为 value, 返回 key 被设置前的旧值

返回: 存在则返回 key 之前的值, 不存在返回 `nil`, key 不为字符串返回 `error`


## STRLEN

格式: `STRLEN key`

说明: 返回键 key 存储的字符串的长度

返回: key 不存在返回 0, 存在返回长度, 不为字符串返回 `error`


## APPEND

格式: `APPEND key value`

说明: key 存在时将 value 追加到原值的末尾; key 不存在时执行和 `SET key value` 一致的操作

返回: 追加 value 后 key 的值的长度


## SETRANGE

格式: `SETRANGE key offset value`

说明: 从 offset 开始覆盖 key 的值, 若字符串长度比偏移量小, 则使用零字节 `\x00` 填充

返回: 修改后字符串的长度


## GETRANGE

格式: `GETRANGE key start end`

说明: 获取 key 的值包含 start 到 end 作为索引的部分. 支持负数, 表示从末尾开始反向计数.
不支持回绕操作, start 到 end 超出实际字符串索引的部分会被忽略.

返回: 字符串指定部分


## INCR

格式: `INCR key`

说明: 为 key 存储的数字值 +1, 支持 int64

返回: 若 key 不存在作为 0; 若 key 不为数字返回 `error`, 为数字返回 key +1 后的值


## INCRBY

格式: `INCRBY key increment`

说明: 为 key 存储的数字值 +increment

返回: 若 key 不存在作为 0; 若 key 不为数字返回 `error`, 为数字返回 key +increment 后的值


## INCRBYFLOAT

格式: `INCRBYFLOAT key increment`

说明: INCRBY 的浮点数版本; 结果至多保留 17 位小数

返回: 同 `INCRBY`


## DECR

格式: `DECR key`

说明: `INCR` 的减法版本

返回: 类似 `INCR`


## DECRBY

格式: `DECRBY key decrement`

说明: `INCRBY` 的减法版本

返回: 类似 `INCRBY`


## MSET

格式: `MSET key1 value1 [key2 value2 ...]`

说明: 同时为多个键设置值, 为原子操作

返回: 总是返回 `OK`


## MSETNX

格式: `MSETNX key1 value1 [key2 value2 ...]`

说明: 仅当所有值都不存在时, 同时为多个键设置值, 为原子操作. 

返回: 成功返回 `1`, 因某个键已存在而未能执行返回 `0`


## MGET

格式: `MGET key1 [key2 key3 ...]`

说明: 返回给定所有键的值, 若有不存在的用 `nil` 表示

返回: 返回一个列表包含所有给定键的值


## TODO

BIT 相关的命令与 STRALGO LCS

