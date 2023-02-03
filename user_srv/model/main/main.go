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
	// Using custom options
	options := &password.Options{
		SaltLen:      16,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: sha512.New,
	}
	salt, encodedPwd := password.Encode(code, options)
	// 最终生成的密码,使用$进行分割,$算法$盐值$密码
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	// 将密码解析出来, 但是要注意一点,passwordInfo的切片长度=4,第一个值是""
	passwordInfo := strings.Split(newPassword, "$")
	check := password.Verify(code, passwordInfo[2], passwordInfo[3], options)
	fmt.Println(check)
	return newPassword
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
