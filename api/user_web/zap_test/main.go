/**
 * @Author: Bruce
 * @Description:
 * @Date: 2023/2/15 1:24 PM
 */

package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	logger, _ := zap.NewProduction() // Production Env
	defer logger.Sync()              // flushes buffer, if any 刷新缓存
	url := "www.baidu.com"
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}
