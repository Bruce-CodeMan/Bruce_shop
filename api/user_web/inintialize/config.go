/**
 * @Author: Bruce
 * @Description: Initialize the yaml config
 * @Date: 2023/2/16 9:07 AM
 */
package inintialize

import (
	"github.com/spf13/viper"

	"Bruce_shop/api/user_web/global"
)

// GetEnvInfo Get the value of the variable which I set
// For example,
// vim ~/.bash_profile, export BRUCE_SHOP=true export PATH=$PATH:$BRUCE_SHOP
// returns the true / false
func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

// InitConfig Initialize the whole project config yaml file
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
