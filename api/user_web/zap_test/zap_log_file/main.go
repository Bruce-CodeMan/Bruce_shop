/**
 * @Author: Bruce
 * @Description:
 * @Date: 2023/2/15 1:42 PM
 */

package main

import "go.uber.org/zap"

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./project.log",
		"stderr",
		"stdout",
	}
	return cfg.Build()
}

func main() {
	logger, err := NewLogger()
	if err != nil {
		panic(err)
	}
	s := logger.Sugar()
	defer s.Sync()
	url := "www.baidu.com"
	s.Info("Failed to fetch URL",
		zap.String("url", url))
}
