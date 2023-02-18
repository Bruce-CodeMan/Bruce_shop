/**
 * @Author: Bruce
 * @Description: 描述
 * @Date: 2023-02-18 12:42
 */

package initialize

import "go.uber.org/zap"

func InitLogger() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
