# pubsub

## PUBLISH

格式: `PUBLISH channel message`

说明: 将信息发送到指定频道

返回: 接收到 message 的订阅者的数量


## SUBSCRIBE

格式: `SUBSCRIBE channel1 [channel2 channel3 ...]`

说明: 订阅给定的一个或多个频道的信息

返回: 接收到的信息.
订阅频道首先返回一个三元组, 分别表示是否订阅成功, 订阅的频道名, 当前已订阅的频道数量.
之后进入消息等待状态, 每次收到消息返回一个三元组, 分别表示类型, 来源频道名, 消息内容.


## PSUBSCRIBE

格式: `PSUBSCRIBE patten1 [patten2 patten3 ...]`

说明: 使用通配符匹配多个频道进行订阅

返回: 同 SUBSCRIBE


## UNSUBSCRIBE

格式: `UNSUBSCRIBE [channel1 [channel2 channel3 ...]]`

说明: 退订给定的频道

返回: 返回一到多个消息, 不同客户端表现不同


## PUNSUBSCRIBE

格式: `PUNSUBSCRIBE [patten1 [patten2 patten3 ...]]`

说明: 退订符合匹配模式的频道

返回: 返回一到多个消息, 不同客户端表现不同


## PUBSUB

格式: `PUBSUB <subcommand> [argument1 [argument2 argument3 ...]]`

说明: 查看订阅与发布系统状态, 有多个子命令
- `PUBSUB CHANNELS [patten]` 返回活跃频道的列表
- `PUBSUB NUMSUB channel1 [channel2 channel3 ...]` 返回给定频道的名称和订阅者数量, 
不含订阅模式的客户端
- `PUBSUB NUMPAT` 返回客户端订阅的所有模式的总和

返回: 见上子命令

