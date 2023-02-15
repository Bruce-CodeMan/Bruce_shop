/**
 * @Author: Bruce
 * @Description:
 * @Date: 2023/2/15 2:04 PM
 */
package router

import (
	"Bruce_shop/api/user_web/api"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", api.GetUserList)
	}
}
