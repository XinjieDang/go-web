package impl

import (
	"com.dxj/app/api/request"
	"com.dxj/app/api/response"
	"com.dxj/app/model"
	"com.dxj/common/utils"
	"com.dxj/global"
)

type UserDao struct {
}

func (dao *UserDao) CreateUser(user *model.User) (err error) {
	// 新增用户初始密码为123456
	user.PassWord = utils.MD5V("123456", "", 0)
	if err := global.DB.Create(user).Error; err != nil {
		return err
	}
	return
}

func (dao *UserDao) FindListAll(users *[]model.User) (err error) {
	result := global.DB.Find(users)
	if result.Error != nil {
		return result.Error
	}
	return
}
func (dao *UserDao) FindUserById(user *model.User, id string) (err error) {
	result := global.DB.First(user, id)
	if result.Error != nil {
		return result.Error
	}
	return
}
func (dao *UserDao) UpdateUser(user *model.User) (err error) {
	result := global.DB.Updates(user)
	if result.Error != nil {
		return result.Error
	}
	return
}
func (dao *UserDao) DeleteUser(id uint64) (err error) {
	result := global.DB.Delete(&model.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return
}
func (dao *UserDao) LoginUser(userLoginReq *request.UserLogin) (*response.User, error) {
	resp := response.User{}
	result := global.DB.Where("name = ?", userLoginReq.UserName).First(&resp)
	if result.Error != nil {
		return nil, result.Error
	}
	return &resp, result.Error
}
