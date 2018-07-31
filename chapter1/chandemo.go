package main

import (
	"time"
	"fmt"
)

func Producer(factor int, out chan<- int) {
	for {
		out <- factor
	}
}

func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 64)
	go Producer(3, ch)
	go Producer(5, ch)

	go Consumer(ch)

	time.Sleep(5* time.Second)
}