# persistence

## SAVE

格式: `SAVE`

说明: 执行同步保存操作, 将当前实例所有数据快照以 RDB 文件的形式保存到硬盘.
以阻塞的方式进行, 一般使用 `BGSAVE` 异步保存.

返回: 成功返回 `OK`


## BGSAVE

格式: `BGSAVE`

说明: 后台异步保存当前数据库的数据到硬盘.

返回: 执行后立即返回 `OK`, 然后 fork 出一个新的子进程进行保存.
可通过 LASTSAVE 命令查看相关信息, 判断是否成功.


## BGREWRITEAOF

格式: `BGREWRITEAOF`

说明: 执行一个 AOF 文件重写操作

返回: 反馈信息


## LASTSAVE

格式: `LASTSAVE`

说明: 最近一次成功保存数据到磁盘的时间

返回: UNIX 时间戳
