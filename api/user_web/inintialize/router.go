/**
 * @Author: Bruce
 * @Description: Initialize the router config
 * @Date: 2023/2/15 2:09 PM
 */
package inintialize

import (
	"github.com/gin-gonic/gin"

	"Bruce_shop/api/user_web/middlewares"
	"Bruce_shop/api/user_web/router"
)

// InitRouters Initialize the router settings
func InitRouters() *gin.Engine {
	Router := gin.Default()
	// 配置跨域请求
	Router.Use(middlewares.Cors())
	ApiRouter := Router.Group("/v1")
	router.InitUserRouter(ApiRouter)
	router.InitBaseRouter(ApiRouter)
	return Router
}
