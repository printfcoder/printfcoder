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
		Addr:     "10.177.98.179:6379",
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
