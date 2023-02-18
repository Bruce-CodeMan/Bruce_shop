/**
 * @Author: Bruce
 * @Description: gRPC config structs connection .yaml
 * @Date: 2023-02-18 12:09
 */

package config

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type ServerConfig struct {
	MysqlInfo MysqlConfig `mapstructure:"mysql" json:"mysql"`
}