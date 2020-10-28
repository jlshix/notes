# debug

## PING

格式: `PING`

说明: 服务器正常运行返回一个 `PONG`

返回: `PONG`


## ECHO

格式: `ECHO message`

说明: 打印 `message`

返回: `message` 本身


## OBJECT

格式: `OBJECT subcommand [arguments1 [arguments2 ...]]`

说明: 从内部查看给定 key 的 Redis 对象, 一般用于调试或了解内部存储方式.
subcommand 包括:
- `OBJECT REFCOUNT key`: key 存储的值的引用计数
- `OBJECT ENCODING key`: key 存储的值的存储方式
- `OBJECT IDLETIME key`: key 的空闲秒数, 即没有被读写的时长
- `OBJECT HELP`: 帮助信息

返回: 见说明


## SLOWLOG

格式: `SLOWLOG subcommand [argument]`

说明: 记录查询执行时间的日志系统, 不包括 IO, 保存在内存中.
subcommand 包括:
- `SLOWLOG SET slowlog-log-slower-than 100`: 记录所有查询时间大于等于100微秒的查询
- `SLOWLOG SET slowlog-max-len 1000`: 最多保存 1000 条 slowlog 日志
- 以上两个命令 `SET` 改为 `GET`, 去掉最后一个参数, 可查询当前设定的值
- `SLOWLOG GET number`: 获取指定数量的 slowlog, 不指定 number 时返回所有日志. 最新的日志
最先被打印. 每项为四元组: 标识符, 时间戳, 执行时长, 完整命令字符串列表.
- `SLOWLOG LEN`: 当前日志数量
- `SLOWLOG RESET`: 清空 SLOWLOG

返回: 见上


## MONITOR

格式: `MONITOR`

说明: 实时打印服务器收到的命令, 调试用.

返回: 每行格式为 `时间戳 [数据库号码 地址:端口] 命令`


## DEBUG OBJECT

格式: `DEBUG OBJECT key`

说明: 不应在客户端使用的调试命令

返回: key 的相关信息, key 不存在返回 error


## DEBUG SEGFAULT

格式: `DEBUG SEGFAULT`

说明: 执行一个不合法的内存访问使 Redis 崩溃, 用于 bug 模拟

返回: 无
