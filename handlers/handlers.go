package handlers

import "github.com/gofiber/fiber/v2"

type AppHandler struct{}

func NewAppHandler() *AppHandler {
	return &AppHandler{}
}

func (a *AppHandler) HandleGetIndex(ctx *fiber.Ctx) error {
	return ctx.Render("index", fiber.Map{})
}
