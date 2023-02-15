/**
 * @Author: Bruce
 * @Description:
 * @Date: 2023/2/15 2:04 PM
 */
package router

import (
	"Bruce_shop/api/user_web/api"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	zap.S().Info("配置用户相关的router")
	{
		UserRouter.GET("list", api.GetUserList)
	}
}
