package responseutil

import "github.com/gofiber/fiber/v2"

type IResponseUtil interface {
	Create(c *fiber.Ctx, message string, data interface{}) error
	Success(c *fiber.Ctx, message string, data interface{}) error

	InternalServer(c *fiber.Ctx, message string) error
	BadRequest(c *fiber.Ctx, message string) error
	Unauthorized(c *fiber.Ctx, message string) error
	NotFound(c *fiber.Ctx, message string) error
}

type ResponseUtil struct{}

func NewResponseUtil() IResponseUtil {
	return &ResponseUtil{}
}

func (u *ResponseUtil) Create(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": message,
		"data":    data,
	})
}

func (u *ResponseUtil) Success(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": message,
		"data":    data,
	})
}

func (u *ResponseUtil) BadRequest(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"status":  fiber.StatusBadRequest,
		"message": message,
	})
}

func (u *ResponseUtil) InternalServer(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status":  fiber.StatusInternalServerError,
		"message": message,
	})
}

func (u *ResponseUtil) Unauthorized(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status":  fiber.StatusUnauthorized,
		"message": message,
	})
}

func (u *ResponseUtil) NotFound(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"status":  fiber.StatusNotFound,
		"message": message,
	})
}
