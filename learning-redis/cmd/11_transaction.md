# transaction

## MULTI

格式: `MULTI`

说明: 标记一个事务块的开始

返回: `OK`


## EXEC

格式: `EXEC`

说明: 执行所有事务块中的命令

返回: 事务块中所有命令的返回值, 按先后顺序排列. 当操作被打断时返回空值 `nil`


## DISCARD

格式: `DISCARD`

说明: 取消事务, 放弃执行事务块内的所有命令

返回: `OK`


## WATCH

格式: `WATCH key1 [key2 key3 ...]`

说明: 监视一到多个 key, 若事务执行前这些 key 的内容有变动则事务将被打断

返回: `OK`


## UNWATCH

格式: `UNWATCH`

说明: 取消 WATCH 命令对所有 key 的监视. 若 WATCH 后执行了 EXEC 或 DISCARD 则无需执行 UNWATCH

返回: `OK`
