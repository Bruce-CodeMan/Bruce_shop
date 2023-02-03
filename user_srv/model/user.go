/*
 * @Date: 2023-02-03 14:08:25
 * @Author: Bruce
 * @Description:
 */
package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int32     `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool `gorm:"column:is_deleted"`
}

/*
	1. 密码,采用md5
	2. 对称加密:[加密和解密是同一把钥匙]
	3. 非对称加密:[加密和解密采用的不是用一把钥匙]
	4. md5:信息摘要算法
*/

type User struct {
	BaseModel
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password string     `gorm:"type:varchar(100);not null"`
	NickName string     `gorm:"type:varchar(20)"`
	Birthday *time.Time `gorm:"type:datetime"`
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment 'female表示女,male表示男'"`
	Role     int        `gorm:"column:role;default:1;type:int comment '1表示普通用户,2表示管理员'"`
}
