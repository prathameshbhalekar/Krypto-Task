package utils

import (
	"net/smtp"

	"github.com/spf13/viper"
)

func SendMail(to []string) error {
	from := viper.GetString("EMAIL_ID")
	password := viper.GetString("EMAIL_PASSWORD")

	// Receiver email address.
	// to := []string{
	// 	"prathameshbhalekar13@gmail.com",
	// }

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("Alert from price tracker!")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	return err
}
