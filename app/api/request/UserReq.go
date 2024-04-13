package request

type UserLogin struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password"  binding:"required"`
}
type CreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      uint8  `json:"age"`
	PassWord string `json:"password"`
}
