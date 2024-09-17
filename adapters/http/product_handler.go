package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/horThana/Backend/core/domain"
	"github.com/horThana/Backend/core/services"
)

type HttpProductHandler struct {
    service services.ProductService
}

func NewHttpProductAdapter(service services.ProductService) *HttpProductHandler{
	return &HttpProductHandler{service: service}
}

func (h *HttpProductHandler) CreateProduct(c* fiber.Ctx) error {
	var product domain.Product
	if err := c.BodyParser(&product); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"ไม่สามารถดึงข้อมูลได้"})

	}

	if err := h.service.CreateProduct(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})

	}

	return c.Status(fiber.StatusCreated).JSON(product)
}

func (h *HttpProductHandler) FindProductByID(c* fiber.Ctx) error {
	id := c.Params("id")
	product, err := h.service.FindProductId(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

func (h *HttpProductHandler) FindAllProducts(c* fiber.Ctx) error {
	product, err := h.service.FindAllProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

func (h *HttpProductHandler) DeleteProduct(c* fiber.Ctx) error {
	id := c.Params("id")
	err := h.service.DeleteProduct(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}

	return c.Status(fiber.StatusNoContent).JSON(nil)
}