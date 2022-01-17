package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	waitByWG()
}

func waitByChannel() {
	c := make(chan bool, 100)
	for i:=0; i<100; i++ {
		go func(i int) {
			fmt.Println(i)
			c <- true
		}(i)
	}

	for i:=0; i<100; i++ {
		<- c
	}
}

func waitBySleep() {
	for i:=0; i<100; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(1 * time.Second)
}

func waitByWG() {
	wg := sync.WaitGroup{}
	wg.Add(1000000)
	for i:=0; i<1000000; i++ {
		go func(i int) {
			println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}