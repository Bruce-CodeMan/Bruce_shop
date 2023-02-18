/**
 * @Author: Bruce
 * @Description: initialize the config ,read the config ingo from the config-*.yaml it determines the env variable
 * @Date: 2023-02-18 12:16
 */

package initialize

import (
	"github.com/spf13/viper"

	"Bruce_shop/srvs/user_srv/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	// 从配置文件中读取相对应的配置
	debug := GetEnvInfo("BRUCE_SHOP")
	configFileName := "srvs/user_srv/config-dev.yaml"
	if debug {
		configFileName = "srvs/user_srv/config-dev.yaml"
	}
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}
}
