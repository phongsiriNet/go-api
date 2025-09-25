package handler

import "github.com/gofiber/fiber/v2"

type IUserHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}
