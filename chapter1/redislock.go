package main

/**
 * 基于redis 的锁实现　
*/


import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

func incr() {

	client := redis.NewClient(
		&redis.Options{
			Addr: "localhost:6379",
			Password: "",
			DB: 0,
		}
	)

	var lockkey = "counter_lock"
	var counterKey = "counter"

	// lock
	resp := client.SexNX(lockkey, 1, time.Second * 5)
	lockSuccess, err := resp.Result()
	
	if err != nil || !lockSuccess {
		fmt.Println(err, "lock result:", lockSuccess)
		return
	}

	// counter ++
	getResp := client.Get(counterKey)
	cntValue, err := getResp.Int64()
	if err == nil {
		cntValue ++
		resp := client.Set(counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			fmt.Println("set value error!")
		}
	}
	fmt.Println("current counter is", cntValue)

	delResp := client.Del(lockkey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
		fmt.Println("unlock success!")
	} else {
		fmt.Println("unlock failed", err)
	}
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}
	wg.Wait()
}