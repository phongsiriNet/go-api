package handler

import (
	"go-api/dto"
	"go-api/pkg/service"
	"go-api/utils/responseutil"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userSrv      service.IUserService
	httpResponse responseutil.IResponseUtil
}

func NewUserHandler(userSrv service.IUserService) IUserHandler {
	return &userHandler{
		userSrv:      userSrv,
		httpResponse: responseutil.NewResponseUtil(),
	}
}

func (u *userHandler) Register(c *fiber.Ctx) error {
	req := dto.UserRequest{}
	if err := c.BodyParser(req); err != nil {
		return u.httpResponse.Errors(c, fiber.StatusBadRequest, "Invalid request body")
	}
	err := u.userSrv.RegisterSVC(&req)
	if err != nil {
		return u.httpResponse.Errors(c, fiber.StatusInternalServerError, err.Error())
	}
	return u.httpResponse.Success(c, "User registed succesfully", nil)
}

func (u *userHandler) Login(c *fiber.Ctx) error {
	req := dto.UserRequest{}
	if err := c.BodyParser(req); err != nil {
		return u.httpResponse.Errors(c, fiber.StatusBadRequest, "Invalid request body")
	}
	jwt, err := u.userSrv.LoginSVC(&req)
	if err != nil {
		return u.httpResponse.Errors(c, fiber.StatusInternalServerError, err.Error())
	}
	return u.httpResponse.Success(c, "Login Successfully", jwt)
}
