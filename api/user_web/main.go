/**
 * @Author: Bruce
 * @Description: main function to run the whole project
 * @Date: 2023/2/15 2:00 PM
 */

package main

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"Bruce_shop/api/user_web/inintialize"
	bruceValidator "Bruce_shop/api/user_web/validator"
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

	// Step2, Register validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", bruceValidator.ValidateMobile)
	}

	zap.S().Info("启动服务器")

	_ = Router.Run()
}
