package lead

import (
	"log"

	"github.com/Dejannnn/FiberCrudProj.git/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLead(c *fiber.Ctx) error {
	db := database.DBConn
	var lead Lead
	leadId := c.Params("id")
	db.Find(&lead, leadId)
	return c.JSON(lead)
}
func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)

	return c.JSON(leads)
}

func NewLead(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)
	err := c.BodyParser(lead)
	if err != nil {
		log.Fatal("Unable to parse lead body")
	}

	db.Create(&lead)

	return c.JSON(lead)
}

func UpadeLead(c *fiber.Ctx) error {
	db := database.DBConn
	leadId := c.Params("id")

	var leadModel Lead
	lead := new(Lead)

	err := c.BodyParser(&lead)
	if err != nil {
		log.Fatal("Unable to parse body %v", err)
	}
	db.First(&leadModel, leadId) // find product with id 1

	db.Model(&leadModel).Update(lead)

	return c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	db := database.DBConn
	leadId := c.Params("id")
	var lead Lead

	db.First(&lead, leadId)

	if lead.Name == "" {
		c.Status(404).SendString("No lead found with ID")
	}
	db.Delete(&lead)
	return c.SendString("Lead succesfully deleted")

}
