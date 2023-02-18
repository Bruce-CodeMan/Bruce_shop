/**
 * @Author: Bruce
 * @Description: initialize the logger using zap package,not to log package
 * @Date: 2023-02-18 12:42
 */

package initialize

import "go.uber.org/zap"

func InitLogger() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
