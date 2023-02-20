/**
 * @Author: Bruce
 * @Description: convert the config yaml to the struct
 * @Date: 2023/2/16 9:03 AM
 */

package config

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Name string `mapstructure:"name"`
}

type ServerConfig struct {
	Name        string        `mapstructure:"name"`
	Port        int           `mapstructure:"port"`
	UserSrvInfo UserSrvConfig `mapstructure:"user_srv"`
	JwtInfo     JWTConfig     `mapstructure:"jwt"`
	AliSmsInfo  AliSmsConfig  `mapstructure:"sms"`
	RedisInfo   RedisConfig   `mapstructure:"redis"`
	ConsulInfo  ConsulConfig  `mapstructure:"consul"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}

type AliSmsConfig struct {
	ApiKet    string `mapstructure:"key"`
	ApiSecret string `mapstructure:"secret"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}
