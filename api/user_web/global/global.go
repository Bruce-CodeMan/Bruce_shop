/**
 * @Author: Bruce
 * @Description: Define some global variables
 * @Date: 2023/2/16 9:47 AM
 */

package global

import (
	ut "github.com/go-playground/universal-translator"

	"Bruce_shop/api/user_web/config"
	"Bruce_shop/api/user_web/proto"
)

var (
	ServerConfig  = &config.ServerConfig{}
	NacosConfig   = &config.NacosConfig{}
	Trans         ut.Translator
	UserSrvClient proto.UserClient
)
