package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/horThana/Backend/core/domain"
	"github.com/horThana/Backend/core/services"
)

type HttpPayment_Handler struct {
		service services.PaymentService

}

func NewHttpPaymentAdapter(service services.PaymentService) *HttpPayment_Handler{
	return &HttpPayment_Handler{service: service}
}

func(h *HttpPayment_Handler) CreatePaymentSubscription(c *fiber.Ctx) error {
	var payment domain.PayMentSubscription
	if err := c.BodyParser(&payment); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"ไม่สามารถดึงข้อมูลได้"})

	}

	if err := h.service.CreatePaymentSubscription(payment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})

	}

	return c.Status(fiber.StatusCreated).JSON(payment)
}

func(h *HttpPayment_Handler) FindPaymentSubscriptionByID(c *fiber.Ctx) error {
	id := c.Params("id")
	payment, err := h.service.FindPaymentSubscriptionId(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(payment)
}

func(h *HttpPayment_Handler) FindAllPaymentSubscriptions(c *fiber.Ctx) error {
	payment, err := h.service.FindAllPaymentSubscriptions()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(payment)
}

func(h *HttpPayment_Handler) DeletePaymentSubscription(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.service.DeletePaymentSubscription(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message":"ลบข้อมูลสำเร็จ"})
}
