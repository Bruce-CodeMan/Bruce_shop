/**
 * @Author: Bruce
 * @Description:
 * @Date: 2023/2/16 9:03 AM
 */
package config

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	Name          string `mapstructure:"name"`
	UserSrvConfig `mapstructure:"user_srv"`
}
