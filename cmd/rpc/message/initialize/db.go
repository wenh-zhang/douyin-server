package initialize

import (
	"douyin/cmd/rpc/message/global"
	"douyin/shared/constant"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	var err error
	mysqlConfig := global.MySQLConfig
	dsn := fmt.Sprintf(constant.MySQLDefaultDSN, mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database)
	global.DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		klog.Fatalf("init gorm failed: %s", err)
	}
}
