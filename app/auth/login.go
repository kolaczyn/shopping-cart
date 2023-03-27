package appAuth

import (
	"github.com/kolaczyn/shopping-cart/repo"
	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) (*UserDto, error) {
	user, err := repo.GetUser(email)
	if err != nil {
		return nil, err
	}

	// TOOD this might be a vulnerability, if we return the error to the user
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, err
	}

	return &UserDto{
		Id:    user.Id,
		Email: user.Email,
		Token: "token",
	}, nil
}
