# Redis

1. RedisObject是通过什么进行标记引用的？介绍一下对象共享的机制和过程。
> 答：RO有一个属性叫refcount，通过这个字段来标记此对象的引用计数，RO被创建时，该属性会被初始化为1，而当RO被引用时，就会+1；使用完对象与取消引用，都会-1，直到等于0时便释放。
  一般而言，命令的返回值、小于REDIS_SHARED_INTEGERS的整数、(待补充)。

2. 有限个小于10000的数字集合（set)，Redis如何存储这些数字？这是什么机制？
> 答：当元素个数小于512个（set-max-intset-entries）时，会放在IntSet（排序后的Int Array）中，每个元素占用16 bits。当个数大于512个时，会升级成HashTable，每个元素会包装在RedisObject中，RO会多花销16bytes的空间，但是它通过指针，与共享值连接起来，故而没有其它内存开销，RO都会放在HashTable中，HT的大小随着集合大小改变。

[32Bit数字的保存](https://newbedev.com/memory-efficient-way-to-store-32-bit-signed-integer-in-redis)


3. 简单介绍一下Redis RDB，RDB的优势、劣势。

> 答：RDB是Redis Database File的缩写，Redis通过RDB来备份、同步数据。RDB的优势是其是二进制文件，恢复起来比AOF日志快得多，压缩后同步传送，效率高。RDB快照生成需要一定时间，节点数据越多越慢，无法秒级实时备份，且因为Fork子进程来生成快照，成本高昂，不能频繁操作。

4. 讲讲Redis 快照生成的几个方式。

RDB生成有几个触发方式:<br/>
**BGSAVE**，fork 进程完成，fork时阻塞<br/>
**SAVE** 不fork，阻塞请求完成:<br/>
**SAVE line**，配置文件中有SAVE选项时，在其声明的值阈触发**BGSAVE**，如'save 60 10000'声明60s内写入10000次时执行SAVE，当。有多行配置时，任一满足都会执行。:<br/>
**SHUTDOWN/TERM**，关机命令或终止命令发生时，都会触发SAVE操作，阻塞所有客户端所有命令。
**SYNC**，在Redis连接到其它Redis节点并发出**SYNC**命令开始复制，主节点会打开执行**BGSAVE**（当前没有在执行或都刚执行完）。
**debug reload**，重新加载会触发**BGSAVE**