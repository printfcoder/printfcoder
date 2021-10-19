# 为什么Kafka这么快

1. 基于日志结构，append-only顺序IO读写
2. 批次传送，client与broker批量收发
3. 批压缩

## 参考资料

[面试官：Kafka 为什么快](https://baijiahao.baidu.com/s?id=1675052145445497191&wfr=spider&for=pc)
[why-kafka-is-so-fast](https://medium.com/swlh/why-kafka-is-so-fast-bde0d987cd03)