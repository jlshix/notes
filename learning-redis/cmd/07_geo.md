# geo

地理位置

## GEOADD

格式: `GEOADD key longitude1 latitude1 member1 [longitude2 latitude2 member2 ...]`

说明: 将一到多个 (经度, 纬度, 名字) 三元组添加到 key 中

返回: 新添加的元素数量


## GEOPOS

格式: `GEOPOS key member1 [member2 member3 ...]`

说明: 查询一到多个 member 的经纬度

返回: 二元组组成的数组, 每项为经纬度, 某项不存在为 `nil`


## GEODIST

格式: `GEODIST key member1 member2 [unit]`

说明: 返回两个给定位置之间的距离. 有一个不存在返回 `nil`.
unit 可选 m(米), km(千米), mi(英里), ft(英尺)

返回: 位置有一个不存在返回 `nil`, 都存在返回一个双精度浮点数


## GEORADIUS

格式: `GEORADIUS key longitude latitude radius unit [WITHCOORD][WITHDIST][WITHHASH] [ASC|DSC][COUNT count]`

说明: 以给定经纬度为中心, 返回 key 包含的元素中与中心距离不超过给定最大距离的所有位置元素

参数:
- unit 可选 m(米), km(千米), mi(英里), ft(英尺)

- WITHDIST: 一并返回位置元素与中心之间的距离, 单位与给定的 unit 一致

- WITHCOORD: 一并返回元素的经纬度

- WITHHASH: 一并返回位置的分值, 实际作用不大

- ASC: 从近到远排序

- DSC: 从远到近排序

- COUNT: 限制返回的数量

返回: 一个列表, 在不指定任何 WITH 选项时为名称列表. 否则为嵌套数组, 每项第一项总是元素名称.
命令中各个 WITH 可不分先后指定, 返回中的先后按先后总是名称, 距离, 哈希值, 坐标.


## GEORADIUSBYMEMBER

格式: `GEORADIUSBYMEMBER key member radius unit [WITHCOORD][WITHDIST][WITHHASH] [ASC|DSC][COUNT count]`

说明: 类似 GEORADIUS, 输入为 member 而非 longitude 和 latitude

返回: 和 GEORADIUS 格式一致


## GEOHASH

格式: `GEOHASH key member1 [member2 member3 ...]`

说明: 返回一到多个元素的哈希

返回: 一个列表, 每项都是一个 geohash

