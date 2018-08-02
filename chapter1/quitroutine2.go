package main

// 改进版本　退出取保完成清理工作

import (
	"fmt"
	"time"
	"sync"
)

func worker(cannel chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <- cannel:
			return
			//退出逻辑
		default:
			// 正常流程
			fmt.Println("hello")
		}
	}
}


func main() {
	ch := make(chan bool)

	var wg sync.WaitGroup
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go worker(ch, &wg)
	}

	time.Sleep(5*time.Second) // 5s
	close(ch)

	wg.Wait()
}