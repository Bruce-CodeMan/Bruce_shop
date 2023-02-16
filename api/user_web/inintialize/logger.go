/**
 * @Author: Bruce
 * @Description: Initialize the logger setting
 * @Date: 2023/2/15 2:39 PM
 */
package inintialize

import (
	"go.uber.org/zap"
)

// InitLogger initialize the logger setting
func InitLogger() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
