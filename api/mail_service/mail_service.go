package mail

import (
	"context"
	"log"

	"github.com/KryptoTask/api/utils"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func EnableMailService() {
	ctx := context.Background()
	addr := viper.GetString("RDB_ADDR")
	password := viper.GetString("RDB_PASS")

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0, // use default DB
	})

	pubsub := rdb.Subscribe(ctx, "mailchannel")

	go func() {
		for {
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				panic(err)
			}

			log.Println(msg.Payload)
			err = utils.SendMail([]string{msg.Payload})
			if err != nil {
				panic(err)
			}

			rdb.Ping(ctx)

		}

	}()
}
