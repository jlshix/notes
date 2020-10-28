# lua scripts


## EVAL

格式: `EVAL script numkeys key1 [key2 key3 ...] arg1 [arg2 arg3 ...]`

说明: script 是一段 lua 脚本; numkeys 是涉及到 key 的数量, arg 是参数.
如 `eval "return {KEYS[1],KEYS[2],ARGV[1],ARGV[2]}" 2 key1 key2 first second`
返回 `["key1", "key2", "first", "second"]`; 需要注意的是:
- KEYS 数组保存命令中传入的 keys
- ARGV 数组保存命令中传入的 args
- 下标从 `1` 开始
- 可使用 `redis.call(*args)` 或 `redis.pcall(*args)` 在脚本中执行 redis 命令
- 脚本的执行是原子性的, 效果等同于事务

返回: 脚本的执行结果


## EVALSHA

格式: `EVALSHA sha1 numkeys key1 [key2 key3 ...] arg1 [arg2 arg3 ...]`

说明: 根据给定的 sha1 执行缓存在服务器中的脚本

返回: 脚本的执行结果


## SCRIPT LOAD

格式: `SCRIPT LOAD script`

说明: 将 script 添加到服务器的脚本缓存中, 并不执行.

返回: script 的 SHA1 校验和


## SCRIPT EXISTS

格式: `SCRIPT EXISTS sha11 [sha12 ...]`

说明: 判断一到多个 sha1 对应的脚本是否已被缓存

返回: 一个列表, 每项 `1` 表示已缓存, `0` 表示未缓存


## SCRIPT FLUSH

格式: `SCRIPT FLUSH`

说明: 清除所有 lua 脚本缓存

返回: `OK`


## SCRIPT KILL

格式: `SCRIPT KILL`

说明: 终止当前正在运行的 LUA 脚本, 当且仅当这个脚本未执行任何写操作时生效

返回: 成功 `OK`, 否则 `error`


