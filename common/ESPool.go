package common

import "github.com/TarsCloud/TarsGo/tars"

// Es连接池，用于处理

var logger = tars.GetLogger("ElasticSearch Pool")

func init() {
	logger.Info("初始化完成！")
}
