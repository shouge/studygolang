package main

import (
    "fmt"
    "sync"
)

func main(){
    var wg sync.WaitGroup
    
    done := make(chan int, 10)
  
    for i := 0; i < cap(done); i++ {
        // 增加等待事件个数
        wg.Add(1)

        go func(i int) {
            fmt.Println("hello world!", i)
            // 完成一个事件
            wg.Done()
            done <- i
        }(i)

    }
    
    // 等待Ｎ个事件完成
    wg.Wait()
}