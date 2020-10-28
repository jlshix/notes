# expire

## EXPIRE

格式: `EXPIRE key seconds`

说明: 为 key 设置以秒为单位的生存时间, 归 0 时自动删除. 过期时间的延迟在 1ms 内

返回: 成功返回 `1`; key 不存在或设置失败返回 `0`


## EXPIREAT

格式: `EXPIREAT key timestamp`

说明: 为 key 设置过期时间, timestamp 为 UNIX 时间戳, 如 `1355292000`

返回: 成功返回 `1`; key 不存在或设置失败返回 `0`


## TTL

格式: `TTL key`

说明: 以秒为单位的 key 的剩余生存时间

返回: key 不存在返回 `-2`; 未设置生存时间返回 `-1`; 否则为剩余的秒数


## PERSIST

格式: `PERSIST key`

说明: 移除 key 的生存时间

返回: 移除成功返回 `1`; 不存在或无生存时间返回 `0`


## PEXPIRE

格式: `PEXPIRE key milliseconds`

说明: EXPIRE 的毫秒版本

返回: 成功返回 `1`; key 不存在或设置失败返回 `0`


## PEXPIREAT

格式: `PEXPIREAT key milliseconds-timestamp`

说明: EXPIREAT 的毫秒版本

返回: 成功返回 `1`; key 不存在或设置失败返回 `0`


## PTTL

格式: `PTTL key`

说明: TTL 的毫秒版本

返回: key 不存在返回 `-2`; 未设置生存时间返回 `-1`; 否则为剩余的毫秒数

