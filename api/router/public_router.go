package router

import (
	"github.com/KryptoTask/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func MountPublicRoutes(c *fiber.App) {
	api := c.Group("api")

	{
		api.Get("/token/new", controllers.GetNewAccessToken)
	}

}
