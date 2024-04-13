package main

import (
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
)
import _ "github.com/swaggo/files"
import (
	_ "com.dxj/docs" //重要
	"com.dxj/initialize"
	"com.dxj/router"
)

// 添加注释以描述 server 信息
// @title           go-web测试API
// @version         1.0
// @description     Gin框架生成CRUD
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8085
// @BasePath  /api/v1
// @securityDefinitions.basic  BasicAuth
func main() {
	//config.InitMysqlConfigInfo()
	initialize.GlobalInit()
	//注册路由
	r := router.SetRouter()
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//启动端口为8085的项目
	r.Run(":8085")
}
