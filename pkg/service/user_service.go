package service

import (
	"go-api/dto"
	"go-api/pkg/model"
	"go-api/pkg/repository"

	"go-api/utils/passwordutil"
	"log"
)

type userService struct {
	userRepo repository.IUser
	auth     IAuth
	passUtil passwordutil.IPassword
}

func NewUserSVC(userRepo repository.IUser, jwtSecret string) IUserService {
	return &userService{
		userRepo: userRepo,
		auth:     NewjwtService(jwtSecret),
		passUtil: passwordutil.NewPasswordUtil(),
	}
}

func (u *userService) RegisterSVC(req *dto.UserRequest) error {

	if req.Role == "" {
		req.Role = "customer"
	}
	hashPassword, err := u.passUtil.HashPassword(req.Password)
	if err != nil {
		return err
	}
	newUser := model.User{
		Email:    req.Email,
		Password: hashPassword,
		Role:     req.Role,
	}
	err = u.userRepo.CreateUser(&newUser)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *userService) LoginSVC(req *dto.UserRequest) (string, error) {
	targetUser, err := u.userRepo.GetUser(req.Email)
	if err != nil {
		return "", err
	}
	match := u.passUtil.CheckPasswordHash(req.Password, targetUser.Password)
	if !match {
		return "", err
	}
	jwt, err := u.auth.CreateJwt(targetUser)
	if err != nil {
		return "", err
	}
	return jwt, err
}
