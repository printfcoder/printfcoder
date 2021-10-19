# Redis

1. RedisObject是通过什么进行标记引用的？介绍一下对象共享的机制和过程。
> 答：RO有一个属性叫refcount，通过这个字段来标记此对象的引用计数，RO被创建时，该属性会被初始化为1，而当RO被引用时，就会+1；使用完对象与取消引用，都会-1，直到等于0时便释放。
  一般而言，命令的返回值、小于REDIS_SHARED_INTEGERS的整数、(待补充)。

2. 有限个小于10000的数字集合（set)，Redis如何存储这些数字？这是什么机制？
> 答：当元素个数小于512个（set-max-intset-entries）时，会放在IntSet（排序后的Int Array）中，每个元素占用16 bits。当个数大于512个时，会升级成HashTable，每个元素会包装在RedisObject中，RO会多花销16bytes的空间，但是它通过指针，与共享值连接起来，故而没有其它内存开销，RO都会放在HashTable中，HT的大小随着集合大小改变。

[32Bit数字的保存](https://newbedev.com/memory-efficient-way-to-store-32-bit-signed-integer-in-redis)