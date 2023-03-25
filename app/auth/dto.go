package appAuth

type RegisterFormDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginFormDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDto struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}
