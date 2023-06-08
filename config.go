package tube

type Config struct {
	Topic []TopicConf //初始化topic对应的消息队列
}

type TopicConf struct {
	Name      string //topic名称
	QueueNums int32  //topic对应的队列数
	QueueLen  int32  //queue的长度
}
