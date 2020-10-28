# hyperloglog

用于统计日活的利器

## PFADD

格式: `PFADD key element1 [element2 element3 ...]`

说明: 将一到多个元素添加到 HyperLogLog key 中

返回: key 发生变化返回 `1`, 否则返回 `0`


## PFCOUNT

格式: `PFCOUNT key1 [key2 key3 ...]`

说明: 返回给定 key 的近似基数, 不存在返回 `0`, 多个 key 时会临时合并

返回: 给定 key 的近似基数, 不存在返回 `0`


## PFMERGE

格式: `PFMERGE destkey sourcekey1 [sourcekey2 sourcekey3 ...]`

说明: 将多个 sourcekey 合并至 destkey

返回: `OK`
