package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sfreiberg/gotwilio"
)

func main() {
	var (
		host = flag.String("host", "", "The Host or IP address the service is listening on")
		port = flag.String("port", "3000", "The IP port number the service is listening on")
	)

	flag.Parse()
	address := fmt.Sprintf("%v:%v", *host, *port)

	router := gin.Default()
	router.GET("/receive", func(c *gin.Context) {
		c.String(http.StatusOK, "Hi this is the TwilioBot listening endpoint!")
	})

	router.Run(address)
}

func SendSMS() {
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
