package dao

import (
	"com.dxj/app/api/request"
	"com.dxj/app/api/response"
	"com.dxj/app/model"
)

type UserDaoRepo interface {
	CreateUser(user *model.User) error
	FindListAll(users *[]model.User) error
	FindUserById(user *model.User, id string) error
	UpdateUser(user *model.User) error
	DeleteUser(id uint64) error
	LoginUser(userLoginReq *request.UserLogin) (*response.User, error)
}
