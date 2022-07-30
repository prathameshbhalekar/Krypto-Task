package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/KryptoTask/api/db"
	"github.com/KryptoTask/api/schemas"
	"github.com/go-redis/redis/v9"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
)

func ConnectWebsocket() {
	addr := viper.GetString("RDB_ADDR")
	password := viper.GetString("RDB_PASS")

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0, // use default DB
	})

	// Connect to websocket
	ctx := context.Background()
	conn, _, err := websocket.DefaultDialer.Dial("wss://stream.binance.com/ws/btcusdt@miniTicker", nil)

	if err != nil {
		log.Panic(err.Error())
	}

	go func() {
		for {
			_, message, readErr := conn.ReadMessage()
			if readErr != nil {
				fmt.Println(readErr)
				return
			}

			var resp schemas.WebsocketResponse
			err = json.Unmarshal(message, &resp)

			if err != nil {
				log.Panic(err.Error())
			}

			c, err := strconv.ParseFloat(resp.C, 32)
			if err != nil {
				log.Panic(err.Error())
			}

			// Get triggered alerts
			var alerts []schemas.Alert

			if err := db.
				GetDB().
				Table("alerts").
				Where("((GREATER_THAN AND ? <= ALERT_VALUE) OR (NOT GREATER_THAN AND ? >= ALERT_VALUE)) AND STATUS='ACTIVE'", c, c).
				Find(&alerts).Error; err != nil {
				log.Panic(err.Error())
			}

			for _, alert := range alerts {
				log.Println(alert)

				// Update alert status

				if err := db.GetDB().Exec(`
				UPDATE ALERTS
				SET STATUS = 'TRIGGERED'
				WHERE ALERT_UUID = ?`,
					alert.AlertUUID).Error; err != nil {
					log.Panic(err.Error())
				}

				// Publish alert
				err := rdb.Publish(ctx, "mailchannel", alert.Email).Err()
				if err != nil {
					panic(err)
				}
			}

		}
	}()
}
