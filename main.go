package main

import (
	"fmt"
	"strings"

	"github.com/KryptoTask/api/cache"
	mail "github.com/KryptoTask/api/mail_service"
	"github.com/KryptoTask/api/router"
	"github.com/KryptoTask/api/utils"
	"github.com/KryptoTask/api/websocket"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func healthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func main() {

	//Set global configuration
	utils.ImportEnv()

	// Init redis
	cache.GetRedis()

	// Create Fiber
	app := fiber.New(fiber.Config{})

	app.Get("/", healthCheck)
	app.Get("/health", healthCheck)

	app.Use(logger.New(logger.Config{Next: func(c *fiber.Ctx) bool {
		return strings.HasPrefix(c.Path(), "api")
	}}))

	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "",
		AllowHeaders: "*",
	}))

	// Mount Routes
	router.MountPublicRoutes(app)
	router.MountPrivateRoutes(app)

	// Get Port
	port := utils.GetPort()

	// Connecting to mailing service
	mail.EnableMailService()

	// Connecting to binance websocket
	websocket.ConnectWebsocket()

	// Start Fiber
	err := app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

}
