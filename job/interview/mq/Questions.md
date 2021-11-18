# MQ Questions

1. 如何保证幂等、顺序、不丢失

> Kafka 顺序性：<br />
  Kafka中并不能保证所有分区无状态消息的顺序性，但是可以保证一个分区中所有消息都是有顺序的，同时一个分区只会有一个消费者，所以可以利用这两个特性达到生产与消费的顺序一致性。比如用户的某些行为消息，我们要保证其前后顺序依次送达后端，那可以讲这些消息推送到一个分区；那如果要多个分区来保持高吞吐怎么办，这个时候就可以利用一些带状态的字段，比如用户id作为Kafka Msg的Key，Kafka的Broker会将这些声明了Key的消息按统一的规则散列到不同的分区，同一个Key只要散列规则没有变就会一直分发到同一个分区，所以只要生产者保证投送消息的顺序性，那消费端就能收到一样顺序的消息。

> Kafka不丢失：<br />
  首先，可以肯定的是没有哪个消息系统可以保证数据不丢失，所以在Kafka中，没有保证说消息不丢失，这也不是一个Kafka系统能保证的事，要业务的接入与消费都全方位配合。我们辩证来看这个问题，从发起、经过、结果三个阶段来讨论。
  从发起阶段，也就是生产者包装好消息时，Kafka提供两个选择：同步消息、异步消息。
  使用同步消息，那就丢弃异步并行发送，损失性能，可以确保消息发出到Broker。
  使用异步消息，可以使用get()获取状态，或者使用回调函数判断发送状态，前者会阻塞当前代码等到get()取得值，所以相当于是同步的方式，在追求性能的非显式多线程发送场景下不提倡使用这个。不过两者都可以拿到状态，所以在发送失败后，可以选择再次发送。
  关于主动发起重试还是由Producer来重试，根据业务来，如果重试时不需要附加重试信息，则可以设计Producer的*retries*来指定自动重试次数。
  到消费阶段.....(待补充)



2. Kafka与rocketMQ集群原理、对比

3. MQ最终一致性，分布式事务问题

4. Kafka Key与Partition的关系。
> 答：消息在发送前会组装成K/V结构，消息通过其中的Partition或Key来确认Partition，当没有显式指名分区时，便会通过计算Key的Hash来分区。但是如果没有Key，也没有Partition，那就会先通过topic查询本地分区表（每隔*topic.metadata.refresh.interval.ms*刷新一次）找出分区，如果分区不存在，则会随机算出分区。

5. 描述Kafka Controller
> 答：Controller是Broker中选举出的控制器，它本身是一个Broker，不过它会负责Partition、副本状态管理，比如重分配Partition，其LeaderSelector会选择新Leader，并更新ISR到ZK，同时通过所有副本更新Leader与ISR。

6. Kafka 的ISR有哪些地方维护
> 答：Controller（被选举的Broker）与Leader。Controller见(KafkaController)[#5]，Leader中有线程检测ISR是否有follower脱离，有则将ISR列表返回ZK。

7. Kafka扩容后如何保证数据仍然在老分区


参考资料：<br />
[Kafka Replica机制](https://www.cnblogs.com/caoweixiong/p/12049462.html)
