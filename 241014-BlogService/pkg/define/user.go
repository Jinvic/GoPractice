package define

type UserRegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
