# 内部命令

## MIGRATE

格式: `MIGRATE host port key destination-db timeout [COPY] [REPLACE]`

说明: 将 key 原子性地从当前实例传送到目标实例的指定数据库上, 阻塞两个实例执行此操作.
过程如下:
- 当前实例对 key 执行 `DUMP` 进行序列化
- 传送到目标实例的数据库
- 目标实例对 key 执行 `RESTORE` 并添加
- 当前实例执行 `DEL` 删除此 key

参数:
- timeout 为数据传输的最大时间
- COPY 不删除当前实例的 key
- REPLACE 替换目标实例上已存在的 key


返回: 成功返回 OK 失败返回相应错误


## DUMP

格式: `DUMP key`

说明: 返回 key 的序列化值

返回: key 的序列化值


## RESTORE

格式: `RESTORE key ttl serialized-value [REPLACE]`

说明: 反序列化数据设为 key, 并设置生存时间为 ttl 毫秒

返回: 成功返回 OK 失败返回 error


## SYNC

格式: `SYNC`

说明: 用于 replication 的内部命令

返回: 序列化数据


## PSYNC

格式: `PSYNC master_run_id` offset

说明: 用于 replication 的内部命令

返回: 序列化数据
