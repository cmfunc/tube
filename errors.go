package tube

import "github.com/pkg/errors"

var ErrBrokerUninit = errors.New("tube's broker haven't init") //tube代理没有初始化
