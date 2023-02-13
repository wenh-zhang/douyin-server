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

// GetNewConn create a new connection
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

// GetConn return existing connection
func GetConn() *gorm.DB {
	return db
}
