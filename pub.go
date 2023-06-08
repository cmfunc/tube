package tube

import (
	"sync/atomic"

	"github.com/pkg/errors"
)

// Pub 向broker中投递消息
func (b *Broker) Pub(topic string, data []byte) error {
	if b == nil {
		return ErrBrokerUninit
	}
	bt, ok := b.GetTopic(topic)
	if !ok {
		logger.Infof("topic:%s haven't init, now start init topic queue's num:%d and queue's len:%s", topic, DefaultTopicQueueNum, DefaultTopicQueueLen)
		b.AddTopic(topic, DefaultTopicQueueNum, DefaultTopicQueueLen)
		bt, ok = b.GetTopic(topic)
		if !ok {
			return errors.Errorf("havn't get topic:%s after auto init", topic)
		}
	}
	ch := bt.ch[bt.writeNext]
	atomic.AddInt32(&bt.writeNext, 1)
	atomic.CompareAndSwapInt32(&bt.writeNext, bt.len, 0)
	// 从channel中获取数据，封装为方法，将消息顺序写入文件
	// 记录文件中数据追加的行数，对文件数据不删除；
	ch <- data
	return nil
}
