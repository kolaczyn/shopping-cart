package appAuth

import "github.com/kolaczyn/shopping-cart/repo"

func Register(email string, password string) (UserDto, error) {
	user, err := repo.CreateUser(email, password)
	if err != nil {
		return UserDto{}, err
	}

	return UserDto{
		Id:    user.Id,
		Email: user.Email,
		// TODO
		Token: "token",
	}, nil
}
