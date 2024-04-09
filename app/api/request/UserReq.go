package request

type UserLogin struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password"  binding:"required"`
}
