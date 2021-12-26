package main

import (
	"fmt"
	"time"
)

/**
队列：
队列长度 10，队列元素类型为 int
生产者：
每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
消费者：
每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
 */

func main() {

	ch := make(chan int, 10)
	defer close(ch)
	ticker := time.NewTicker(time.Second)
	// producer
	go func() {
		for _ = range ticker.C {
			ch <- int(time.Now().Unix())
		}
	}()
	// consumer
	for _ = range ticker.C {
		v:= <- ch
		fmt.Println(v)
	}
}
