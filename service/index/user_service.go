package index

import "github.com/xufwind95/go-web-base/model"

type UserService struct {
	UserName string
	Password string
}

func (userService *UserService) AddUser() error {
	userModel := model.UserTestModel{
		Username: userService.UserName,
		Password: userService.Password,
	}
	return model.CreateUser(&userModel)
}

func FindUser(username string) (*UserService, error) {
	userModel := model.UserTestModel{
		Username: username,
	}

	ret, err := model.FindOneUser(&userModel)
	if err != nil {
		return nil, err
	}

	return &UserService{UserName: ret.Username}, nil
}
