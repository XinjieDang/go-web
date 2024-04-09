package response

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      uint8  `json:"age"`
	PassWord string `json:"password"`
	Token    string `json:"token"`
}
