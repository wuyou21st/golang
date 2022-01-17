package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go lock()
	go rlock()
	go wlock()
	time.Sleep(5 * time.Second)
}

func lock() {
	// 互斥锁
	lock := sync.Mutex{}
	for i:=0; i<3; i++ {
		lock.Lock()
		fmt.Println("lock", i)
		lock.Unlock()
	}
}

func rlock() {
	// 读写分离锁
	lock := sync.RWMutex{}
	for i:=0; i<3; i++ {
		lock.RLock()
		defer lock.RUnlock()
		fmt.Println("rlock", i)
	}
}

func wlock() {
	// 读写分离锁
	lock := sync.RWMutex{}
	for i:=0; i<3; i++ {
		lock.Lock()
		fmt.Println("wlock", i)
		lock.Unlock()
	}
}

