/*
 * @Date: 2023-02-03 14:08:25
 * @Author: Bruce
 * @Description:
 */
package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"Bruce_shop/user_srv/model"
)

// genMd5 Generate MD5
func genMd5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}

func main() {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/shop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢查询，阈值
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 这行代码可以将用户表设置成user
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(&model.User{})
}
