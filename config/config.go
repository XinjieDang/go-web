package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Mysql struct {
	Url      string
	Port     int
	Username string
	Password string
	Dbname   string
}

type Redis struct {
	Host     string
	Port     int
	Password string
	DataBase int `mapstructure:"data_base"`
}
type Jwt struct {
	Admin JwtOption
	User  JwtOption
}

type JwtOption struct {
	Secret string
	TTL    string
	Name   string
}

// 整合配置文件
type Config struct {
	Mysql Mysql `json:"mysql"`
	Redis Redis `json:"redis"`
	Jwt   Jwt
}

func InitMysqlConfigInfo() *Config {
	dataBytes, err := os.ReadFile("./config/application.yaml")
	if err != nil {
		fmt.Println("读取文件失败：", err)
	}
	//fmt.Println("yaml 文件的内容: \n", string(dataBytes))
	var config = Config{}
	//配置文件的内容转化为结构体
	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		fmt.Println("解析 yaml 文件失败：", err)
	}
	//fmt.Printf("config → %+v\n", config) // config → {Mysql:{Url:127.0.0.1 Port:3306} Redis:{Host:127.0.0.1 Port:6379}}
	//初始化并连接数据库
	mp := make(map[string]any, 2)
	//配置文件内容转换到map
	err = yaml.Unmarshal(dataBytes, mp)
	if err != nil {
		fmt.Println("解析 yaml 文件失败：", err)
	}
	return &config
	//fmt.Printf("===================map====================== → %+v", config) // config → {Mysql:{Url:127.0.0.1 Port:3306} Redis:{Host:127.0.0.1 Port:6379}}
}
