package controllers

import (
	"github.com/KryptoTask/api/utils"
	"github.com/gofiber/fiber/v2"
)

func GetNewAccessToken(c *fiber.Ctx) error {
	// Generate a new Access token.
	token, err := utils.GenerateNewAccessToken()
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(fiber.Map{
		"success":      true,
		"access_token": token,
	})
}
