package tube

import (
	"os"
)

// 数据持久化，硬盘顺序读写
// 按天分文件存储消息持久化数据
func WriteDiskFile(file string, data []byte) {
	os.WriteFile(file, data, os.ModeAppend)
}
