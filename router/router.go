package router

import (
	"com.dxj/app/api"
	"com.dxj/middle"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	//登录公共路由
	LoginGroup := r.Group("login")
	{
		controller := api.UserController{}
		LoginGroup.POST("/userLogin", controller.Login)
	}
	/**
	  用户User路由组
	*/
	userGroup := r.Group("user")
	//私有路由使用JWT验证
	userGroup.Use(middle.VerifyJWTAdmin())
	{
		//增加用户User
		controller := api.UserController{}
		userGroup.POST("/users", controller.Create)
		userGroup.PUT("/users", controller.Update)
		userGroup.GET("/users", controller.FindList)
		userGroup.GET("/users/:id", controller.FindById)
		userGroup.DELETE("/users/:id", controller.Delete)
	}
	return r
}
