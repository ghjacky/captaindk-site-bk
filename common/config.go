package common

import (
	"fmt"
	"github.com/spf13/viper"
	"net"
	"os"
)

// 服务配置struct定义
type SConfig struct {
	Listen       string   // 服务监听地址，例：127.0.0.1:8080
	LogFile      *os.File // 日志文件
	SMysqlConfig          // mysql配置
	FileStore    string   // 文件存储位置
}

// 全局配置对象定义
var (
	Config     = &SConfig{}
	DefaultLog = os.Stdout
)

const (
	DefaultListen        = "127.0.0.1:8080"
	DefaultMysqlBind     = "127.0.0.1:3306"
	DefaultFileStore     = "/tmp"
	DefaultMysqlDB       = "cdks"
	DefaultMysqlUser     = "root"
	DefaultMysqlPassword = "roothjack"
)

// 从配置文件读取具体配置信息（赋值给Config以使全局使用）
func ReadConfigurationFromFile(file string) {
	// 读取配置文件
	Log.Infof("读取配置文件：%s", file)
	viper.SetConfigFile(file)
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("配置文件读取错误: %s", err.Error()))
	}
	// 读取成功，检测配置项
	validateOrSetDefaultConfiguration(Config)
	return
}

// 逐一检测配置项，可用则继续，不存在或为空则置默认值，不可用则panic
func validateOrSetDefaultConfiguration(config *SConfig) {
	// 日志
	if len(viper.GetString("main.log")) == 0 {
		Log.Warnf("配置项：main.log 缺失，使用默认值：%s", DefaultLog.Name())
		config.LogFile = DefaultLog
	} else {
		var err error
		if config.LogFile, err = os.OpenFile(viper.GetString("main.log"), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644); err != nil {
			Log.Warnf("日志文件打开失败：%s，使用默认值：%s", err.Error(), DefaultLog.Name())
			config.LogFile = DefaultLog
		}
	}
	// 服务监听地址
	config.Listen = valideNetAddr("main.listen", DefaultListen)
	config.FileStore = valideString("main.file_store", DefaultFileStore)
	// 数据库
	config.SMysqlConfig.Bind = valideNetAddr("mysql.bind", DefaultMysqlBind)
	config.SMysqlConfig.Database = valideString("mysql.database", DefaultMysqlDB)
	config.SMysqlConfig.User = valideString("mysql.user", DefaultMysqlUser)
	config.SMysqlConfig.Password = valideString("mysql.password", DefaultMysqlPassword)
	config.SMysqlConfig.ConnParams = "charset=utf8&parseTime=True&loc=Local"
}

// 网络地址配置检测
func valideNetAddr(configPattern, defaultConfig string) string {
	var addr = viper.GetString(configPattern)
	if len(addr) == 0 {
		Log.Warnf("配置项：%s 缺失，使用默认值：%s", configPattern, defaultConfig)
		return defaultConfig
	} else if _, err := net.ResolveTCPAddr("tcp4", addr); err != nil {
		panic(fmt.Sprintf("配置项：%s 有误：%s", configPattern, err.Error()))
	}
	return addr
}

// 字符串配置检测
func valideString(configPattern, defaultConfig string) string {
	var s = viper.GetString(configPattern)
	if len(s) == 0 {
		return defaultConfig
	}
	return s
}
