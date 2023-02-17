/**
 * @Author: Bruce
 * @Description: convert the config yaml to the struct
 * @Date: 2023/2/16 9:03 AM
 */
package config

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	Name        string        `mapstructure:"name"`
	UserSrvInfo UserSrvConfig `mapstructure:"user_srv"`
	JwtInfo     JWTConfig     `mapstructure:"jwt"`
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
