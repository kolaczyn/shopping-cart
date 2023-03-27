package repo

func GetUser(email string) (*User, error) {
	var user User
	err := getDb().Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func CreateUser(email string, password string) (*User, error) {
	passwordHash, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Email:        email,
		PasswordHash: passwordHash,
	}

	err = getDb().Create(&user).Error
	if err != nil {
		return nil, err
	}
	user, err = GetUser(email)
	return user, err
}
