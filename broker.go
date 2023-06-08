package tube

import (
	"sync"
)

// Broker 实际消息队列的代理
// 优化点：控制锁的粒度
type Broker struct {
	//同步锁，增删topic时需要锁住
	lock *sync.RWMutex
	// 存放数据的队列
	topic map[string]*BrokerTopic
	// 生产消息的生产者(应该放在queue中处理)

	// 消费数据的消费者(应该放在queue中处理)

}

type BrokerTopic struct {
	ch        []chan []byte
	len       int32
	writeNext int32
	readNext  int32
}

// AddTopic 动态添加topic
// 优化点：等待broker中生产者消费者空闲时，添加
func (b *Broker) AddTopic(name string, queueNums int32, queueLen int32) {
	b.lock.Lock()
	defer b.lock.Unlock()
	queues := make([]chan []byte, queueNums)
	for i := 0; i < int(queueNums); i++ {
		queues[i] = make(chan []byte, queueLen)
	}
	b.topic[name] = &BrokerTopic{
		ch:        queues,
		len:       queueLen,
		writeNext: queueNums,
	}
}

// GetTopic 获取topic
// 降低锁的粒度
func (b *Broker) GetTopic(name string) (*BrokerTopic, bool) {
	b.lock.RLock()
	defer b.lock.RUnlock()
	bt, ok := b.topic[name]
	if !ok {
		return nil, false
	}
	return bt, true
}

func NewBroker(conf *Config) *Broker {
	if conf == nil {
		panic("conf is nil")
	}
	b := &Broker{
		lock:  &sync.RWMutex{},
		topic: map[string]*BrokerTopic{},
	}
	for _, topicConf := range conf.Topic {
		queues := make([]chan []byte, topicConf.QueueNums)
		for i := 0; i < int(topicConf.QueueNums); i++ {
			queues[i] = make(chan []byte, topicConf.QueueLen)
		}
		b.topic[topicConf.Name] = &BrokerTopic{
			ch:        queues,
			len:       topicConf.QueueLen,
			writeNext: topicConf.QueueNums,
		}
	}
	return b
}
