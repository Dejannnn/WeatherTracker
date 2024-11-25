package router

import (
	"github.com/Dejannnn/FiberCrudProj.git/lead"
	"github.com/gofiber/fiber/v2"
)

func Router() *fiber.App {
	app := fiber.New()

	v1 := app.Group("/api/v1")

	v1.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("Fuck")
	})
	v1.Get("/lead", lead.GetLeads)
	v1.Get("/lead/:id", lead.GetLead)
	v1.Put("/lead/:id", lead.UpadeLead)
	v1.Post("/lead", lead.NewLead)
	v1.Delete("/lead/:id", lead.DeleteLead)

	return app
}
