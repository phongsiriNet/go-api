package examplesvc

import (
	"go-api/pkg/repository/examplerepo"
	"go-api/utils/responseutil"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IExampleService interface {
	Example() (string, *responseutil.HttpResponse)
}

type ExampleService struct {
	ExampleRepository examplerepo.IExampleRepository
	HttpResponse      responseutil.IResponseUtil
}

func NewExampleServiced(dbconn *gorm.DB) IExampleService {
	return &ExampleService{
		ExampleRepository: examplerepo.NewExampleRepository(dbconn),
		HttpResponse:      responseutil.NewResponseUtil(),
	}
}

func (s *ExampleService) Example() (string, *responseutil.HttpResponse) {
	example := s.ExampleRepository.Example()
	if example == "" {
		return "", s.HttpResponse.HttpResponse(fiber.StatusNotFound, "Example not found")
	}
	return example, nil
}
