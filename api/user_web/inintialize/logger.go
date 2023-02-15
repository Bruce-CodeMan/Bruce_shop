/**
 * @Author: Bruce
 * @Description:
 * @Date: 2023/2/15 2:39 PM
 */
package inintialize

import (
	"go.uber.org/zap"
)

func InitLogger() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
