package main

import (
	"content-grpc/core"
	"content-grpc/global"
	"content-grpc/initialize"
)

func main() {

	global.VIPER = core.Viper() 		//初始化viper
	global.LOG = core.Zap()				//初始化zap日志库
	global.DB = initialize.Gorm()   	//初始化GORM连接数据库
	initialize.Redis() 					//初始化redis连接
	global.GRPC = core.InitGrpcServer(global.CONFIG.GRPC.Port) //初始化GRPC连接

	defer global.GRPC.Stop()
}

