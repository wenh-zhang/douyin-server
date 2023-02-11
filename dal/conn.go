package dal

import (
	"douyin/pkg/constant"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = GetNewConn()
}

func GetNewConn() *gorm.DB {
	db, err := gorm.Open(mysql.Open(constant.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	return db
}

func GetConn() *gorm.DB {
	return db
}
