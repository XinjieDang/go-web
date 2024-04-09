package main

import (
	"com.dxj/initialize"
	"com.dxj/router"
)

func main() {
	//config.InitMysqlConfigInfo()
	initialize.GlobalInit()
	//注册路由
	r := router.SetRouter()
	//启动端口为8085的项目
	r.Run(":8085")
}
