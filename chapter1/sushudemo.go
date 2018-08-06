package main 

import "fmt"

func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}


func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	
	go func() {
		for{
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()

	return out
}

func main() {
	ch := GenerateNatural() 
	
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime)
	}
}


/**
* 基于select 实现管道的超时判断
*
*/

/*
select {
case v := <-in:
	fmt.Println(v)
case <-time.After(time.Second):
	return //超时
}
*/

/**
* 基于select 实现管道的超时判断
*
*/

/*
select {
case v := <-in:
	fmt.Println(v)
default:
	// 没有数据
}
*/