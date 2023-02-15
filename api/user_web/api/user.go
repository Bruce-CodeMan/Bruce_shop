/**
 * @Author: Bruce
 * @Description:
 * @Date: 2023/2/15 2:03 PM
 */

package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetUserList(ctx *gin.Context) {
	zap.S().Info("获取用户列表页")
}
