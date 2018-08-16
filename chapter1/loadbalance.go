package main

import (
	"time"
	"fmt"
	"math/rand"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	// 更新随机种子
}

var endpoints = []string {
	"192.168.1.1:8080",
	"192.168.1.２:8080",
	"192.168.1.３:8080",
	"192.168.1.４:8080",
	"192.168.1.５:8080",
	"192.168.1.６:8080",
	"192.168.1.７:8080",
	"192.168.1.８:8080",
	"192.168.1.９:8080",
}

// 索引数组洗牌
func shuffle(slice []int) {
	// 这种随机缺少随机种子　洗牌不均匀
	/*
	for i := 0; i ＜ len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	} 
	*/

	// 经典的fisher-yetas算法　随机一个值　放在数组末尾 ｎ-1个元素的数组中再随机放在末尾 
	
	for i := len(indexex); i > 0; i-- {
		lastIdx := i - 1
		idx := rand.Int(i)
		indexex[lastIdx], indexex[idx] = indexex[idx], indexex[lastIdx]
	}

	// rand.Perm() 实现了该算法

}

func request(params map[string]interface{}) error {
	var indexex = []int {0,1,2,3,4,5,6,7,8}
	var err error

	shuffle(indexex)
	maxRetryTimer := 3

	idx := 0
	for i:=0; i < maxRetryTimer: i++ {
		// 
		err = apiRequest(params, indexex[idx])
		if err == nil {
			break
		}
		idx ++
	}

	if err != nil {
		return err
	}
	
	return nil
}