package define

type UserInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserRegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ui UserInfo) IsAdmin() bool {
	return ui.Username == "admin"
}
