# config

## CONFIG SET

格式: `CONFIG SET parameter value`

说明: 动态调整服务器配置而无需重启

返回: 设置成功返回 `OK`, 否则返回一个错误


## CONFIG GET

格式: `CONFIG GET parameter`

说明: 获取运行中 Redis 的参数信息, 可使用 `*` 通配符

返回: 给定配置参数的值


## CONFIG RESETSTAT

格式: `CONFIG RESETSTAT`

说明: 重置 INFO 命令中的某些统计数据

返回: `OK`


## CONFIG REWRITE

格式: `CONFIG REWRITE`

说明: 将服务器当前配置重新写入到启动服务器的配置文件中.

返回: `OK`, 文件不存在或写入失败则返回一个 `error`
