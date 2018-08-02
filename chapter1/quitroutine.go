// 通过select 跟default 可以实现一个Goroutine的退出控制
package main

import (
	"fmt"
	"time"
)


func worker(cannel chan bool) {
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
	go worker(ch)

	time.Sleep(3*time.Second) // 5s
	ch <- false
}