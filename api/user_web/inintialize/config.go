/**
 * @Author: Bruce
 * @Description:
 * @Date: 2023/2/16 9:07 AM
 */
package inintialize

import (
	"Bruce_shop/api/user_web/global"
	"github.com/spf13/viper"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	debug := GetEnvInfo("BRUCE_SHOP")
	configFileName := "api/user_web/config-dev.yaml"
	if debug {
		configFileName = "api/user_web/config-proc.yaml"
	}
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}
}
