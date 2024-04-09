package initialize

import (
	"com.dxj/config"
	"com.dxj/global"
)

func GlobalInit() {
	// 配置文件初始化
	global.Config = config.InitMysqlConfigInfo()
	// Log初始化
	//global.Log = logger.NewMySlog(global.Config.Log.Level, global.Config.Log.FilePath)
	// Gorm初始化
	global.DB = InitDatabase(global.Config)
	// Redis初始化
	global.Redis = initRedis()
	// Router初始化
	//router := routerInit()
	//return router
}
