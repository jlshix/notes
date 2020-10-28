# bitmap

## SETBIT

格式: `SETBIT key offset value`

说明: 对 key 存储的字符串值, 设置或清除指定偏移量上的 bit. offset 大于字符串位长度时
会进行伸展, 空白位置以 0 填充.

返回: 指定偏移量原来存储的 bit


## GETBIT

格式: `GETBIT key offset`

说明: 获取 key 指定偏移量上的 bit

返回: 返回 key 指定偏移量上的 bit, key 不存在或 offset 大于位长度时返回 `0`


## BITCOUNT

格式: `BITCOUNT key [start] [end]`

说明: 给定字符串区间中 1 的数量

返回: `1` 的数量


## BITPOS

格式: `BITPOS key bit [start] [end]`

说明: key 中第一个值为 bit 的位置

返回: 位置整数, 找不到为 `-1`


## BITOP

格式: `BITOP operation destkey key1 [key2 key3 ...]`

说明: 对一到多个 key 进行位操作, operation 可以是 `AND OR NOT XOR`, 结果保存到 destkey.
key 长度不同时补 0; `NOT` 仅限一个 key, 其他都可以是多个 key.

返回: destkey 的长度, 与所有 key 中最长的相等


## BITFIELD

格式: `BITFIELD key [GET type offset] [SET type offset value][INCRBY type offset increment] [OVERFLOW WRAP|SAT|FAIL]`

说明: BITFIELD 将字符串作为二进制位组成的数组进行操作, 拥有四个子命令
- `GET type offset` 返回指定的二进制位范围
- `SET type offset value` 对指定的二进制位范围进行设置, 返回旧值
- `INCRBY type offset increment` 对指定的二进制位进行假发操作, 返回旧值
- `OVERFLOW WRAP|SAT|FAIL` INCRBY 溢出时的行为, 默认为 WRAP, 命令无返回
    - WRAP: 回绕, 即取模, 舍去最高位, 即 int8 的 127 + 1 为 -128
    - SAT: 饱和计算, 超过最大值取最大值, 小于最小值取最小值
    - FAIL: 拒绝执行, 不更改原数据, 结果为空值

返回: 一个数组, 对应子命令的结果
