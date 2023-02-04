/*
 * @Date: 2023-02-03 14:08:25
 * @Author: Bruce
 * @Description:
 */
package main

import (
	"Bruce_shop/user_srv/model"
	"crypto/sha512"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// genMd5 Generate MD5
func genMd5(code string) string {
	options := &password.Options{
		SaltLen:      16,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: sha512.New,
	}
	salt, encodedPwd := password.Encode(code, options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	passwordInfo := strings.Split(newPassword, "$")
	check := password.Verify(code, passwordInfo[2], passwordInfo[3], options)
	fmt.Println(check)
	return newPassword
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/shop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
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

	for i := 0; i < 10; i++ {
		user := model.User{
			NickName: fmt.Sprintf("Bruce-%d", i),
			Mobile:   fmt.Sprintf("1394082924%d", i),
			Password: genMd5("admin123"),
		}
		db.Save(&user)
	}

	// _ = db.AutoMigrate(&model.User{})
}
