package tube

import (
	"sync/atomic"

	"github.com/pkg/errors"
)

func (b *Broker) Sub(topic string, fn func(data []byte) error) error {
	if b == nil {
		return ErrBrokerUninit
	}
	bt, ok := b.GetTopic(topic)
	if !ok {
		return errors.Errorf("topic:%s haven't init, now start init topic queue's num:%d and queue's len:%d", topic, DefaultTopicQueueNum, DefaultTopicQueueLen)
	}
	// 轮询从queues中获取消息内容
	ch := bt.ch[bt.readNext]
	atomic.AddInt32(&bt.readNext, 1)
	atomic.CompareAndSwapInt32(&bt.readNext, bt.len, 0)
	data := <-ch
	return fn(data)
}
