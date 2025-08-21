package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/x01000101x/ticket-project/models"
)

type TicketHandler struct {
	repository models.TicketRepository
}

func NewTicketHandler(router fiber.Router, repository models.TicketRepository) {
	handler := &TicketHandler{
		repository: repository,
	}

	router.Get("/", handler.GetMany)
	router.Get("/:eventId", handler.GetOne)
	router.Post("/", handler.CreateOne)
	router.Post("/validate", handler.ValidateOne)
}

func (h *TicketHandler) GetMany(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	tickets, err := h.repository.GetMany(context)

	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "",
		"data":    tickets,
	})
}

func (h *TicketHandler) GetOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	eventId, _ := strconv.Atoi(ctx.Params("eventId"))

	ticket, err := h.repository.GetOne(context, uint(eventId))

	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "",
		"data":    ticket,
	})

}

func (h *TicketHandler) CreateOne(ctx *fiber.Ctx) error {
	ticket := &models.Ticket{}

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(ticket); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	ticket, err := h.repository.CreateOne(context, ticket)

	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "succes create data!",
		"data":    ticket,
	})
}

func (h *TicketHandler) ValidateOne(ctx *fiber.Ctx) error {
	validateBody := &models.ValidateTicket{}

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(validateBody); err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	validateData := make(map[string]interface{})
	validateData["entered"] = true

	ticket, err := h.repository.UpdateOne(context, validateBody.TicketId, validateData)

	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Welcome to the show!",
		"data":    ticket,
	})

}
