package controllers

import (
	"injection-go/entities"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type membershipHandler struct {
	memberService entities.MemberUsecaseRepository
}

func NewMembershipHandler(memberService entities.MemberUsecaseRepository) *membershipHandler {
	return &membershipHandler{
		memberService: memberService,
	}
}

func (h *membershipHandler) GetAll(c *fiber.Ctx) error {
	// Get all members from usecase
	memberServices, err := h.memberService.GetAll()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(memberServices)
}

func (h *membershipHandler) GetById(c *fiber.Ctx) error {
	// Get ID from request
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Get member by ID from usecase
	memberService, err := h.memberService.GetById(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(memberService)
}
