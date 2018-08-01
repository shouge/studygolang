package pubsub

import (
	"sync"
	"time"
)

type (
	subscriber chan interface{} //　订阅者为一个管道
	topicFunc func(v interface{}) bool // 主题为一个过滤器
)

type Publisher struct {
	m sync.RWMutex // 读写锁
	buffer int // 订阅队列的缓存大小
	timeout time.Duration　//　发布超时时间
	subsrcibers map[subscriber]topicFunc // 订阅者信息
}

// 构建发布者对象
func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher {
		buffer: buffer,
		timeout: publishTimeout,
		subsrcibers: make(map[subscriber]topicFunc),
	}
}

// 添加新的订阅
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

// 添加新的订阅
func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subsrcibers[ch] = topic
	p.m.Unlock()
	return ch
}

// 退出一个订阅
func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()
	
	delete(p.subsrcibers, sub)
	close(sub)
}

// 发布一个主题
func (p *Publisher) Publish(v interface{}) {
	p.m.Lock()
	defer p.m.Unlock()
	
	var wg sync.WaitGroup
	for sub, topic := range p.subsrcibers {
		wg.add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}

	wg.Wait()
}

// 关闭发布者对象　同时关闭所有的订阅者管道
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subsrcibers {
		delete(p.subsrcibers, sub)
		close(sub)
	}
}

// 发送主题
func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	// 这里先检查不为nil 再进行调用
	if topic != nil && !topic(v) {
		return
	}

	// 过滤之后　通过管道传递给订阅者
	select {
	case sub <- v:
	case <- time.After(p.timeout):
	}

}

