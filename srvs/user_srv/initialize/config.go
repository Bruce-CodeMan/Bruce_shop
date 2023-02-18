/**
 * @Author: Bruce
 * @Description: 描述
 * @Date: 2023-02-18 12:16
 */

package initialize

import "github.com/spf13/viper"

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	// 从配置文件中读取相对应的配置
	debug := GetEnvInfo("BRUCE_SHOP")
	configFileName := "/srvs/user_srv/config-dev.yaml"
	if debug {
		configFileName = "/srvs/user_srv/config-dev.yaml"
	}
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	
}
