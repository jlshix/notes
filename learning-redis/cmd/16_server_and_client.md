# server and client

## AUTH

格式: `AUTH password`

说明: 验证密码; 前提是在配置文件中配置了 `requirepass` 或执行了
`CONFIG SET requirepass password`.

返回: 匹配返回 `OK`, 否则返回一个 `error`


## QUIT

格式: `QUIT`

说明: 请求服务器关闭与当前客户端的连接

返回: `OK` 但不显示, 因为已经退出


## INFO

格式: `INFO [section]`

说明: 以一种易于转换且易于阅读的格式返回 Redis 服务器的各种信息和统计数值.
section 的可选值为:
- server: 服务器的信息
- clients: 已连接客户端信息
- memory: 服务器的内存信息
- persistence: RDB 与 AOF 持久化的信息
- status: 一般统计信息
- replication: 复制集信息
- all: 所有信息
- default 默认信息

不带 section 时默认为 default

返回: 服务器信息


## SHUTDOWN

格式: `SHUTDOWN [SAVE | NOSAVE]`

说明: 执行 SHUTDOWN 后服务器会执行以下操作:
- 停止所有客户端
- 如果只是有一个保存点在等待, 执行 SAVE 命令
- 如果 AOF 选项已打开, 更新 AOF 文件
- 关闭 redis server

可选 SAVE 和 NOSAVE, 强制保存或强制不保存, 覆盖服务器原有的设置

返回: 失败返回 `error`, 成功不返回任何消息, 客户端自动退出


## TIME

格式: `TIME`

说明: 返回当前服务器的时间

返回: 一个二元组, 分别为秒数和微妙数


## CLIENT GETNAME

格式: `CLIENT GETNAME`

说明: 返回 `CLIENT SETNAME` 为此连接设置的名字, 未设置返回 `nil`

返回: 连接名称


## CLIENT KILL

格式: `CLIENT KILL ip:port`

说明: 关闭地址为 `ip:port` 的客户端. 由于 Redis 是单线程, 被断开的客户端若有命令正在执行
则会执行完毕后关闭连接.

返回: 指定客户端存在且关闭成功返回 `OK`


## CLIENT LIST

格式: `CLIENT LIST`

说明: 以人类可读的方式返回所有连接到服务器的客户端信息和统计数据

返回: 多行字符串, 每行以 `key1=value1 key2=value2` 的方式展示


## CLIENT SETNAME

格式: `CLIENT SETNAME connection-name`

说明: 为当前连接分配一个名字

返回: 设置成功返回 `OK`
