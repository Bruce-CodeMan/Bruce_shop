/**
 * @Author: Bruce
 * @Description: main function to run the whole project
 * @Date: 2023/2/15 2:00 PM
 */

package main

import (
	"Bruce_shop/api/user_web/utils"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"Bruce_shop/api/user_web/global"
	"Bruce_shop/api/user_web/inintialize"
	bruceValidator "Bruce_shop/api/user_web/validator"
)

func main() {
	// Step 1, initialize the logger
	// 		   initialize the router
	// 		   initialize the config
	// 		   initialize the translation
	//     	   initialize the srv_conn
	inintialize.InitLogger()
	inintialize.InitConfig()
	Router := inintialize.InitRouters()
	if err := inintialize.InitTrans("en"); err != nil {
		panic(err)
	}
	inintialize.InitSrvConn()

	port, err := utils.GetFreePort()
	if err != nil {
		zap.S().Info("获取动态端口号出错")
		return
	}
	global.ServerConfig.Port = port
	// Step 2, Register validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", bruceValidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}
	_ = Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port))
}
