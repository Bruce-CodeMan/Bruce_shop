/**
 * @Author: Bruce
 * @Description:
 * @Date: 2023/2/17 5:42 PM
 */
package router

import (
	"Bruce_shop/api/user_web/api"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("captcha", api.GetCaptcha)
		BaseRouter.POST("send_sms", api.SendSms)
	}
}
