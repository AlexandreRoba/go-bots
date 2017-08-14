package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sfreiberg/gotwilio"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accountSid := os.Getenv("TWILIO_ACCOUNTSID")
	authToken := os.Getenv("TWILIO_TOKEN")
	phoneNumber := os.Getenv("TWILIO_NUMBER")

	client := gotwilio.NewTwilioClient(accountSid, authToken)
	smsResponse, exception, err := client.SendSMS(phoneNumber, "+32473293080", "SMS Sent from GO", "", "")
	if err != nil {
		log.Fatal("Error Sending SMS:", err, exception)
	}
	log.Println(smsResponse)
}
