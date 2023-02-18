/**
 * @Author: Bruce
 * @Description: global variables using in the system
 * @Date: 2023-02-18 12:11
 */

package global

import (
	"Bruce_shop/srvs/user_srv/config"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
)
