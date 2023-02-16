/**
 * @Author: Bruce
 * @Description:
 * @Date: 2023/2/16 9:07 AM
 */
package inintialize

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"Bruce_shop/api/user_web/global"
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
	fmt.Println(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}

	fmt.Println(global.ServerConfig.Name)
	fmt.Println(global.ServerConfig.Host)
	fmt.Println(global.ServerConfig.Port)
	zap.S().Infof("配置信息:&v", global.ServerConfig)
}
