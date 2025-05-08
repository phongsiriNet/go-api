package examplehandler

import (
	"go-api/pkg/service/examplesvc"
	"go-api/utils/responseutil"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IExampleHandler interface {
	Example(c *fiber.Ctx) error
}

type ExampleHandler struct {
	ExampleService examplesvc.IExampleService
	Response       responseutil.IResponseUtil
}

func NewExampleHandler(dbconn *gorm.DB) IExampleHandler {
	return &ExampleHandler{
		ExampleService: examplesvc.NewExampleServiced(dbconn),
		Response:       responseutil.NewResponseUtil(),
	}
}

func (h *ExampleHandler) Example(c *fiber.Ctx) error {
	example, err := h.ExampleService.Example()
	if err != nil {
		return h.Response.Errors(c, err.Status, err.Message)
	}
	return h.Response.Success(c, "success", example)
}
