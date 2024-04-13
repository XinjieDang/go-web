package service

import (
	"com.dxj/app/api/request"
	"com.dxj/app/api/response"
	"com.dxj/app/dao/impl"
	"com.dxj/app/model"
	"com.dxj/common/e"
	"com.dxj/common/utils"
	"com.dxj/global"
	"fmt"
	"log"
	"time"
)

/*
*
新建User信息
*/
type UserService struct {
}

type IUserService interface {
	CreateUser(user *model.User) error
	FindList(users *[]model.User) error
	FindById(user *model.User, id string) error
	UpdateUser(user *model.User) error
	DeleteUser(id uint64) error
	UserLogin(userLoginReq *request.UserLogin) (*response.User, error)
}

var dao = impl.UserDao{}

func (service *UserService) CreateUser(user *model.User) (err error) {
	err = dao.CreateUser(user)
	if err != nil {
		return err
	}
	return
}
func (service *UserService) FindList(users *[]model.User) (err error) {
	result := dao.FindListAll(users)
	return result
}
func (service *UserService) FindById(user *model.User, id string) (err error) {
	result := dao.FindUserById(user, id)
	return result
}

func (service *UserService) UpdateUser(user *model.User) (err error) {
	result := dao.UpdateUser(user)
	return result
}
func (service *UserService) DeleteUser(id uint64) (err error) {
	result := dao.DeleteUser(id)
	return result
}
func (service *UserService) UserLogin(userLoginReq *request.UserLogin) (*response.User, error) {
	//查询用户是否存在
	currUser, err := dao.LoginUser(userLoginReq)
	fmt.Printf("%+v\n", currUser)
	if err != nil || currUser == nil {
		return nil, e.Error_ACCOUNT_NOT_FOUND
	}
	// 2.校验密码
	password := utils.MD5V(userLoginReq.Password, "", 0)
	if password != currUser.PassWord {
		return nil, e.Error_PASSWORD_ERROR
	}
	//校验状态
	// 生成Token
	jwtConfig := global.Config.Jwt.Admin
	var useId uint64 = 16
	token, err := utils.GenerateToken(useId, jwtConfig.Name, jwtConfig.Secret)

	if err != nil {
		return nil, err
	}
	// 4.构造返回数据
	resp := response.User{
		Name:  currUser.Name,
		Token: token,
	}
	//将token 存入redis
	rcw := utils.NewRedisClientWrapper(global.Redis)
	err = rcw.Set("login:token", token, 30*time.Minute)
	if err != nil {
		log.Fatalf("Failed to set key: %v", err)
	}

	return &resp, nil
}
