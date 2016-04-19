package main

import (
	"log"
	"net/http"
	"time"

	twilio "github.com/carlosdp/twiliogo"
)

func main() {
	accountSid := "<PUT-HERE>"
	authToken := "<PUT-HERE>"

	client := twilio.NewClient(accountSid, authToken)

	for {
		time.Sleep(5 * time.Second)

		resp, err := http.Get("http://127.0.0.1:10080/get")
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode == http.StatusNotFound {
			continue
		}

		if resp.StatusCode != http.StatusOK {
			log.Fatalf("invalid http code %d", resp.StatusCode)
		}

		message, err := twilio.NewMessage(client, "{{ twilio number }}", "{{ your number }}", twilio.Body("Alert!"))
		if err != nil {
			log.Fatal(err)
		}
		log.Print(message.Status)
	}
}
