package app

func Login(email string, password string) (UserDto, error) {
	// TODO implement
	user := UserDto{
		Id:    1,
		Email: email,
		Token: "token",
	}
	return user, nil
}
