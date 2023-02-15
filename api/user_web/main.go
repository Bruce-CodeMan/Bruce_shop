/**
 * @Author: Bruce
 * @Description:
 * @Date: 2023/2/15 2:00 PM
 */

package main

import (
	"Bruce_shop/api/user_web/inintialize"
	"go.uber.org/zap"
)

func main() {
	// Step 1, initialize the logger
	// Step 1 ,initialize the router
	inintialize.InitLogger()
	Router := inintialize.InitRouters()

	zap.S().Info("启动服务器")

	Router.Run()
}
