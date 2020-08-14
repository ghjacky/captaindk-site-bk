package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type SMysqlConfig struct {
	Bind       string `json:"bind"`
	Database   string `json:"database"`
	User       string `json:"user"`
	Password   string `json:"password"`
	ConnParams string `json:"conn_params"`
}

var Mysql = &gorm.DB{}

func InitMysql() {
	Log.Infoln("初始化mysql连接")
	var err error
	Mysql, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		Config.SMysqlConfig.User, Config.SMysqlConfig.Password, Config.SMysqlConfig.Bind,
		Config.SMysqlConfig.Database, Config.SMysqlConfig.ConnParams))
	if err != nil {
		Log.Fatalf("无法连接到mysql服务，程序退出! %s", err.Error())
	} else {
		Mysql = Mysql.LogMode(true)
	}
	Log.Infoln("mysql连接初始化完成")
}
