package handler

import (
	"go-api/dto"
	"go-api/pkg/service"
	"go-api/utils/responseutil"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type productHandler struct {
	proSrv       service.IProductService
	HttpResponse responseutil.IResponseUtil
}

func NewProductHandler(proSrv service.IProductService) IProductHandler {
	return &productHandler{
		proSrv:       proSrv,
		HttpResponse: responseutil.NewResponseUtil(),
	}
}

func (h *productHandler) CreateProduct(c *fiber.Ctx) error {
	req := new(dto.ProductRequest)
	if err := c.BodyParser(req); err != nil {
		return h.HttpResponse.Errors(c, fiber.StatusBadRequest, "Invalid request body")
	}
	product, err := h.proSrv.CreateProductSVC(req)
	if err != nil {
		return h.HttpResponse.Errors(c, fiber.StatusInternalServerError, err.Error())
	}

	return h.HttpResponse.Create(c, "Product created Succesfully", product)
}

func (h *productHandler) GetProducts(c *fiber.Ctx) error {
	products, err := h.proSrv.GetProducts()
	if err != nil {
		return h.HttpResponse.Errors(c, fiber.StatusInternalServerError, err.Error())
	}

	return h.HttpResponse.Success(c, "Product queried succesfully", products)
}

func (h *productHandler) GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		return h.HttpResponse.Errors(c, fiber.StatusBadRequest, "Invalid product ID")
	}

	product, err := h.proSrv.GetProduct(uint(productID))
	if err != nil {
		return h.HttpResponse.Errors(c, fiber.StatusInternalServerError, err.Error())
	}
	return h.HttpResponse.Success(c, "Product queried succesfully", product)

}

func (h *productHandler) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		return h.HttpResponse.Errors(c, fiber.StatusBadRequest, "Invalid product ID")
	}

	req := new(dto.ProductRequest)
	if err = c.BodyParser(req); err != nil {
		return h.HttpResponse.Errors(c, fiber.StatusBadRequest, "Invalid request body")
	}
	updatedProduct, err := h.proSrv.UpdateProductSVC(uint(productID), req)

	if err != nil {
		return h.HttpResponse.Errors(c, fiber.StatusInternalServerError, err.Error())
	}

	return h.HttpResponse.Success(c, "Product is updated successfully", updatedProduct)
}

func (h *productHandler) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		return h.HttpResponse.Errors(c, fiber.StatusBadRequest, "Invalid product ID")
	}

	err = h.proSrv.DeleteProductSVC(uint(productID))
	if err != nil {
		return h.HttpResponse.Errors(c, fiber.StatusInternalServerError, err.Error())
	}
	return h.HttpResponse.Success(c, "Product is deleted succesfully", nil)
}

//==================================== send json fiber ==============================================
// func (u *ResponseUtil) Success(c *fiber.Ctx, message string, data interface{}) error {
//     return c.Status(fiber.StatusOK).JSON(fiber.Map{
//         "message": message,
//         "data":    data,
//     })
// }

// resp := HttpResponse{Status: 200, Message: "Deleted successfully"}
// return c.Status(resp.Status).JSON(resp)
