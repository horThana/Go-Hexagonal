package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/horThana/Backend/core/domain"
	"github.com/horThana/Backend/core/services"
)
type HttpUserHandler struct {
		service services.UserService
}

func NewHttpUserAdapter(service services.UserService) *HttpUserHandler{
	return &HttpUserHandler{service: service}
	
}

//CreateUser is a method
func(h *HttpUserHandler) CreateUser(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"cannot parse body"})
	}
	if err := h.service.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

//FindUSerByID is a method
func(h *HttpUserHandler) FindUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.service.FindUseById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

//FindAllUser is a method
func(h *HttpUserHandler) FindAllUsers(c *fiber.Ctx) error {
	user, err := h.service.FindAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

//DeleteUser is a method
func(h *HttpUserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.service.DeleteUser(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message":"User deleted"})
}
