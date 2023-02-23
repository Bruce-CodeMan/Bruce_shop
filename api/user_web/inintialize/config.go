/**
 * @Author: Bruce
 * @Description: Initialize the yaml config
 * @Date: 2023/2/16 9:07 AM
 */

package inintialize

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"

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
	if err := v.Unmarshal(global.NacosConfig); err != nil {
		panic(err)
	}

	// read config information from nacos
	sc := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port:   uint64(global.NacosConfig.Port),
		},
	}

	// 创建clientConfig
	cc := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.NameSpace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "api/user_web/tmp/nacos/log", //去掉tmp前面的/，这样就会默认保存到当前项目目录下
		CacheDir:            "api/user_web/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group})

	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(content), &global.ServerConfig)
	if err != nil {
		zap.S().Fatal("读取配置信息失败: ", err.Error())
	}

}
