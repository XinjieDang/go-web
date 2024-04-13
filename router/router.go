package router

import (
	"com.dxj/app/api"
	"com.dxj/middle"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{ //登录公共路由
		LoginGroup := v1.Group("login")
		{
			controller := api.UserController{}
			LoginGroup.POST("/userLogin", controller.Login)

		}
		/**
		  用户User路由组
		*/
		userGroup := v1.Group("user")
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
	}
	return r
}
