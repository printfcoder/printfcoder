# MQ Questions

1. 如何保证幂等、顺序、不丢失

2. Kafka与rocketMQ集群原理、对比

3. MQ最终一致性，分布式事务问题

4. Kafka Key与Partition的关系。
> 答：消息在发送前会组装成K/V结构，消息通过其中的Partition或Key来确认Partition，当没有显式指名分区时，便会通过计算Key的Hash来分区。但是如果没有Key，也没有Partition，那就会先通过topic查询本地分区表（每隔*topic.metadata.refresh.interval.ms*刷新一次）找出分区，如果分区不存在，则会随机算出分区。


5. Kafka扩容后如何保证老数据仍然在老分区
