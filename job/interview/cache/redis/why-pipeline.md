Redis服务基于TCP编写，其基于请求响应模型（客户端Req/服务端Rsp）提供TCP服务，一般而言，一个Redis请求通常由如下几步构成：

1. 客户端创建socket并发送请求到服务端，然后读取该socket，通常该过程是阻塞的，直到这个服务端响应完成
2. 服务端处理该请求中的命令，然后发送响应到客户端

我们使用一个包含4条递增指令的序列请求来描述这个模型：

> Client: INCR X
> Server: 1
> Client: INCR X
> Server: 2
> Client: INCR X
> Server: 3
> Client: INCR X
> Server: 4

客户端与服务端通过网络连接，可能是本地环回网络，也或者是广域网间隔数跳的主机之间，不论网络的延迟如何，请求数据包都需要时间从客户端发送到服务端，然后数据库包再带着响应从服务端回到客户端。

这个过程所耗的时间称为**RTT**（Round Trip Time，也叫RTD，Round-Trip Delay），由此，一串连接的请求中，网络延迟带来的性能损耗是显而易见的。如果RTT是250毫秒，哪怕服务端一次能接触100K的请求，我们也只能1s中请求4次，如果放在环回网络中，可能RTT只需要0.014-0.017毫秒，客户端1秒也只能完成70余次请求，也会有很多损耗在网络通信上，需要100K/70台客户端才能用完服务端的能力。

为了避免这种情况，Redis提供了Pipelining管道能力。

## Pipelining

既然Server端已经是支持Req/Rsp模型的了，那它与客户端便是解耦的，也即是说它是可以在客户端收到老请求的响应前处理下一个请求的，所以客户端其实是可以一次投送多个指令到服务端而不用等到前一个处理完成后再投送下一个，然后也一次性读取所有响应的结果。

管道技术在早几十年前已经实现，POP3协议也有类似实现来下载多个邮件。

我们开头的示例便可以通过Pipelining改造成如下：

> Client: INCR X
> Client: INCR X
> Client: INCR X
> Client: INCR X
> Server: 1
> Server: 2
> Server: 3
> Server: 4

这样我们就不用像之前一样花费4次请求，仅需要一次即可。

> 服务端处理Pipelining批量指令时，是将其强制放到队列中处理的，所以，如果有大批量的指令，客户端应该把它们切成多个小批次，每个批次放置合理的指令，比如10K个指令，读取与响应完成后再投送下一批10K个请求。每个批次处理的速度大体是相等的，不过需要为这10K队列分配额外的内存。

## 优势

Pipelining并不只是降低了因为网络延迟带来的性能损耗，它同时也提升了服务端处理请求的效率。

如果不使用Pipelining，从访问缓存的数据结构与包装响应来看，处理每个指令本身损耗很低，但是从socket IO的角度看，就很浪费。原因在于syscall，socket IO需要调用系统的read()和write()的方法，这里在操作系统中会有一个用户域到内核域的转换，上下文切换造成巨大的速度损耗。

在Pipelining中，多个指令是复用一个read()，而多个响应也复用一个write()，随着指令数量增加，Pipelining带来的性能呈10x增长。

### 示例

我们使用代码来验证Pipelining在实际大批量查询中的效果与普通查询的对比。

```golang
package main

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	withPipelining()
	withoutPipelining()
}

func newClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}

func withoutPipelining() {
	cli := newClient()
	defer cli.Close()

	start := time.Now()

	ctx := context.Background()
	for i := 0; i < 10000; i++ {
		cli.Get(ctx, "redis:pipelining:test")
	}

	elapsed := time.Since(start)
	log.Printf("withoutPipelining took %s", elapsed)
}

func withPipelining() {
	cli := newClient()
	defer cli.Close()

	start := time.Now()

	p := cli.Pipeline()
	ctx := context.Background()
	for j := 0; j < 10; j++ {
		for i := 0; i < 1000; i++ {
			p.Get(ctx, "redis:pipelining:test")
		}

		p.Exec(ctx)
	}
	elapsed := time.Since(start)
	log.Printf("withPipelining took %s", elapsed)
}
```

测试机器的Redis在另一机房，Ping到该机房的延迟如下，约3.3ms

```shell
PING 10.x.x.x (10.x.x.x) 56(84) bytes of data.
64 bytes from 10.x.x.x: icmp_seq=1 ttl=128 time=3.41 ms
64 bytes from 10.x.x.x: icmp_seq=2 ttl=128 time=3.29 ms
```

我们请求10000次，可以估算出普通查询大概为10000*3.3=33000ms，约33s，我们看看实际运行效果

```shell
2021/05/26 09:30:14 withPipelining took 53.988457ms
2021/05/26 09:30:46 withoutPipelining took 32.159512029s
```

在常规延迟下，10000条指令Pipelining是普通查询的约600倍。

## 总结

Pipelining适用于一次批量查询，在实时性要求不高的业务，或一次业务事务中有多个分散的keys要查询的可以组合在一起，使用Pipelining批量提交。使用时要切记key数不能过多，最好做一次压测，看业务数据结构在实际线上网络中可以批量的大小是多少，因为在Server端，Pipelining是基于队列方式查询的，一次提交太多指令可能会造成队列阻塞，造成延迟，同时也要考虑队列的容量对内存大小的影响。

## 扩展阅读

1. [RTT](https://en.wikipedia.org/wiki/Round-trip_delay)
2. [Redis是长连接还是短连接](./short-or-long-connection.md)
3. [Redis Script VS Pipelining](./script-vs-pipelining.md)
4. [用户域与内核域](../os/userland-vs-kernal.md)