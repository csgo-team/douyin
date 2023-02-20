package DB

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db gorm.DB

//func GetDB() *gorm.DB {
//	//这里open里面的设置我的数据库账户密码的设置
//	db, err := gorm.Open("mysql", fmt.Sprintf("%v:%v@/ByteDance?charset=utf8mb4&parseTime=True&loc=Local", config.DataBaseName, config.Password))
//	if err != nil {
//		println("数据库初始化出错", err)
//		log.Fatal(err)
//	}
//	return db
//}
func GetDB() *gorm.DB {
	//这里open里面的设置我的数据库账户密码的设置

	db, err := gorm.Open(mysql.Open("root:123456@/ByteDance?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		PrepareStmt:            true, //缓存预编译命令
		SkipDefaultTransaction: true, //禁用默认事务操作
	})
	if err != nil {
		println("数据库初始化出错", err)
		log.Fatal(err)

	}
	return db
}
