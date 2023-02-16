/**
 * @Author: Bruce
 * @Description: main function to run the whole project
 * @Date: 2023/2/15 2:00 PM
 */

package main

import (
	"go.uber.org/zap"

	"Bruce_shop/api/user_web/inintialize"
)

func main() {
	// Step 1, initialize the logger
	// 		   initialize the router
	// 		   initialize the config
	// 		   initialize the translation
	inintialize.InitLogger()
	inintialize.InitConfig()
	Router := inintialize.InitRouters()
	if err := inintialize.InitTrans("en"); err != nil {
		panic(err)
	}

	zap.S().Info("启动服务器")

	Router.Run()
}
