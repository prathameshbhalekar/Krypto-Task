package router

import (
	"github.com/KryptoTask/api/controllers"
	"github.com/KryptoTask/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func MountPrivateRoutes(c *fiber.App) {
	alerts := c.Group("api/alerts")

	{
		alerts.Post("/create", middleware.JWTProtected(), controllers.CreateAlert)
		alerts.Get("/fetch", middleware.JWTProtected(), controllers.GetAlerts)
		alerts.Delete("/delete/:id", middleware.JWTProtected(), controllers.DeleteAlert)

	}

}
