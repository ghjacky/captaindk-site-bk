package main

import (
	"captaindk.site/common"
	"captaindk.site/model"
	"captaindk.site/router"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 捕获系统信号
func catchOsSignal() {
	sig := make(chan os.Signal, 0)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	_ = <-sig
	common.Log.Infof("服务中断")
	os.Exit(0)
}

func main() {
	// 从运行参数中读取配置文件路径
	configPath := flag.String("config", "./configs/config.toml", "指定配置文件路径，默认为: ./configs/config.toml")
	flag.Parse()
	// 初始化配置
	common.ReadConfigurationFromFile(*configPath)
	// 初始化日志
	common.InitLog()
	defer func() {
		_ = common.Config.LogFile.Close()
	}()
	// 初始化数据库
	common.InitMysql()
	common.Mysql.AutoMigrate(&model.SArticle{}, &model.SCategory{}, &model.STag{}, &model.SFile{})
	defer func() {
		_ = common.Mysql.Close()
	}()
	// 初始化路由
	router.RegisterRouter()

	// 服务启动
	go catchOsSignal()
	if err := router.Router.Run(common.Config.Listen); err != nil {
		panic(fmt.Sprintf("服务启动失败：%s", err.Error()))
	}
}
