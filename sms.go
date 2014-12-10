package ws

import (
	"errors"
	"github.com/sfreiberg/gotwilio"
	"os"
)

func SMS(to string, message string) error {
	accountSid := os.Getenv("TWILIO_SID")
	authToken := os.Getenv("TWILIO_TOKEN")
	from := os.Getenv("TWILIO_NUMBER")
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	_, exception, _ := twilio.SendSMS(from, to, message, "", "")

	if exception != nil {
		return errors.New(exception.Message)
	} else {
		return nil
	}
}
