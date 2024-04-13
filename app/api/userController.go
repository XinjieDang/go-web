package api

import (
	"com.dxj/app/api/request"
	"com.dxj/app/model"
	"com.dxj/app/service"
	"com.dxj/common"
	"com.dxj/common/e"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var userService = service.UserService{}

type UserController struct {
}

// GetTodo
// @Summary 添加用户
// @Description 添加用户
// @Tags users
// @Accept json
// @Produce json
// @Param user body request.CreateUser true "用户实体"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router       /users/ [post]
func (UserController *UserController) Create(c *gin.Context) {
	//定义一个User变量
	var user model.User

	//将调用后端的request请求中的body数据根据json格式解析到User结构变量中
	c.BindJSON(&user)
	//将被转换的user变量传给service层的CreateUser方法，进行User的新建
	err := userService.CreateUser(&user)
	//判断是否异常，无异常则返回包含200和更新数据的信息
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, common.Result{
			Code: e.SUCCESS,
			Data: user,
		})
	}
}

// GetTodo
// @Summary 查询全部用户信息
// @Description 查询全部用户信息
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router       /user/users [get]
func (UserController *UserController) FindList(c *gin.Context) {
	//定义一个User变量
	var users []model.User
	result := userService.FindList(&users)
	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error()})
	} else {
		c.JSON(http.StatusOK, common.Result{
			Code: e.SUCCESS,
			Msg:  e.GetMsg(e.SUCCESS),
			Data: users,
		})
	}
}

// GetTodo
// @Summary 查询用户信息通过id
// @Description 查询用户信息通过id
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "用户id"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router       /user/users/{id} [get]
func (UserController *UserController) FindById(c *gin.Context) {
	var user model.User
	id := c.Param("id")
	result := userService.FindById(&user, id)
	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error()})
	} else {
		c.JSON(http.StatusOK, common.Result{
			Code: e.SUCCESS,
			Msg:  e.GetMsg(e.SUCCESS),
			Data: user,
		})
	}
}
func (UserController *UserController) Update(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	result := userService.UpdateUser(&user)
	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error()})
	} else {
		c.JSON(http.StatusOK, common.Result{
			Code: e.SUCCESS,
			Msg:  e.GetMsg(e.SUCCESS),
			Data: user,
		})
	}
}
func (UserController *UserController) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	result := userService.DeleteUser(id)
	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error()})
	} else {
		c.JSON(http.StatusOK, common.Result{
			Code: e.SUCCESS,
			Msg:  e.GetMsg(e.SUCCESS),
			Data: result,
		})
	}
}

// GetTodo
// @Summary 用户登录
// @Description 用户登录
// @Tags users
// @Accept json
// @Produce json
// @Param userLoginReq body request.UserLogin true "登录用户信息"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router       /login/userLogin [post]
func (UserController *UserController) Login(c *gin.Context) {
	var userLoginReq request.UserLogin
	c.BindJSON(&userLoginReq)
	resp, result := userService.UserLogin(&userLoginReq)
	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error()})
	} else {
		c.JSON(http.StatusOK, common.Result{
			Code: e.SUCCESS,
			Msg:  e.GetMsg(e.SUCCESS),
			Data: resp,
		})
	}

}
