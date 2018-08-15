package main

import (
	"sync"
	"fmt"
)

type Lock struct {
	c chan struct{}
}

func NewLock() Lock {
	var l Lock
	// 管道的缓存为1
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}

// GetLock
func (l Lock) Lock() bool {
	lockResult := false
	select {
	case <-l.c:
		lockResult = true
	default:
	}
	return lockResult
}

func (l Lock) Unlock() {
	l.c <- struct{}{}
}

var counter int 

func main() {
	var lock = NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			// 没有获取锁
			if !lock.Lock() {
				fmt.Println("lock failed.")
				return
			}
			counter ++ 
			fmt.Println("current counter", counter)
			lock.Unlock()
		}()
	}

	wg.Wait()
}


/**
 * 单机测试有问题
 *
 */