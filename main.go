package main

import (
	"fmt"
	"gin_blog/dao"
	"gin_blog/gb"
	"gin_blog/routers"
	"go.uber.org/zap"
)

func init() {
	gb.LogConf()
}

func main() {
	gb.Logger.Info("start project...")
	err := dao.InitMySQL()
	if err != nil {
		fmt.Printf("init MySQL failed, err:%v\n", err)
		gb.Logger.Error("init MySQL failed", zap.Any("error", err))
		return
	}
	r := routers.SetupRouter() // 初始化路由
	r.Run(":9999")
}
