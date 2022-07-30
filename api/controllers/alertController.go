package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KryptoTask/api/cache"
	"github.com/KryptoTask/api/db"
	"github.com/KryptoTask/api/schemas"
	"github.com/KryptoTask/api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func CreateAlert(ctx *fiber.Ctx) error {
	body := new(schemas.Alert)

	// Extract metadata for user uuid
	metadata, err := utils.ExtractTokenMetadata(ctx)

	if err != nil {
		return ctx.Status(400).JSON(err)
	}

	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(400).JSON(err)
	}

	// Getting current market price
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=USD&order=market_cap_desc&per_page=1&page=1&sparkline=false")
	if err != nil {
		return ctx.Status(500).JSON(err)
	}

	defer resp.Body.Close()

	var updates []schemas.CoinGeckoResponse
	err = json.NewDecoder(resp.Body).Decode(&updates)
	if err != nil {
		return ctx.Status(500).JSON(err)
	}

	currentPrice := updates[0].CurrentPrice
	var newAlert schemas.Alert

	if err := db.GetDB().Raw(`
		INSERT INTO ALERTS(
			ALERT_UUID,
			EMAIL,
			GREATER_THAN,
			ALERT_VALUE,
			USER_UUID,
			STATUS
		) 
		VALUES (?, ?, ?, ?, ?, ?)
		RETURNING *`,
		uuid.New().String(),
		body.Email,
		body.AlertValue < currentPrice,
		body.AlertValue,
		metadata.Userid,
		"ACTIVE",
	).Scan(&newAlert).Error; err != nil {
		return ctx.Status(400).JSON(err)
	}

	return ctx.JSON(fiber.Map{
		"result":  newAlert,
		"success": true,
	})
}

func DeleteAlert(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := db.GetDB().
		Table("alerts").
		Where("ALERT_UUID = ?", id).
		Delete(&schemas.Alert{}).
		Error; err != nil {
		return ctx.Status(400).JSON(err)
	}

	return ctx.JSON(fiber.Map{
		"success": true,
	})
}

func GetAlerts(ctx *fiber.Ctx) error {
	status := ctx.Query("status", "All")
	page_str := ctx.Query("page", "0")
	page_size_str := ctx.Query("page_size", "0")
	page, err := strconv.Atoi(page_str)

	if err != nil {
		return ctx.Status(400).JSON(err)
	}

	page_size, err := strconv.Atoi(page_size_str)
	if err != nil {
		return ctx.Status(400).JSON(err)
	}

	//Extract user data from jwt
	metadata, err := utils.ExtractTokenMetadata(ctx)

	if err != nil {
		return ctx.Status(400).JSON(err)
	}

	// Get Cache
	key := schemas.CacheKey{
		UserUUID: metadata.Userid,
		PageNo:   page,
		PageSize: page_size,
		Status:   status,
	}
	key_json, err := json.Marshal(key)
	if err != nil {
		return ctx.Status(500).JSON(err)
	}
	val_json, err := cache.GetValue(string(key_json))

	if err != nil {
		return ctx.Status(500).JSON(err)
	}

	// Return cache if exists
	if val_json != "" {
		var cachedAlerts []schemas.Alert
		err := json.Unmarshal([]byte(val_json), &cachedAlerts)

		if err != nil {
			return ctx.Status(500).JSON(err)
		}

		return ctx.JSON(fiber.Map{
			"success": true,
			"result":  cachedAlerts,
		})

	}

	offset := page * page_size
	var alerts []schemas.Alert
	if status == "All" {
		// If no filter applied
		if err := db.GetDB().
			Table("alerts").
			Where("USER_UUID = ?", metadata.Userid).
			Offset(offset).
			Limit(page_size).
			Find(&alerts).
			Error; err != nil {
			return ctx.Status(400).JSON(err)
		}
	} else {
		// If filter applied
		if err := db.GetDB().
			Table("alerts").
			Where("USER_UUID = ? AND STATUS = ?", metadata.Userid, status).
			Offset(offset).
			Limit(page_size).
			Find(&alerts).
			Error; err != nil {
			return ctx.Status(400).JSON(err)
		}
	}

	new_val_json, err := json.Marshal(alerts)

	if err != nil {
		return ctx.Status(500).JSON(err)
	}

	cache.SetValue(string(key_json), string(new_val_json), viper.GetDuration("CACHE_TIMEOUT"))

	return ctx.JSON(fiber.Map{
		"success": true,
		"result":  alerts,
	})
}
