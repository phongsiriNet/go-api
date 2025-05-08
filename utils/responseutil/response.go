package responseutil

import "github.com/gofiber/fiber/v2"

type IResponseUtil interface {
	HttpResponse(status int, message string) *HttpResponse

	Create(c *fiber.Ctx, message string, data interface{}) error
	Success(c *fiber.Ctx, message string, data interface{}) error

	Errors(c *fiber.Ctx, statusCode int, message string) error
}

type ResponseUtil struct{}

func NewResponseUtil() IResponseUtil {
	return &ResponseUtil{}
}

type HttpResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (u *ResponseUtil) HttpResponse(status int, message string) *HttpResponse {
	return &HttpResponse{
		Status:  status,
		Message: message,
	}
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

func (u *ResponseUtil) Errors(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  statusCode,
		"message": message,
	})
}
