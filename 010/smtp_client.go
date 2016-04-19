package main

import (
	// Standrdowa paczka do obsługi formatowania (wejście/wyjście)
	"fmt"
	"log"
	// Paczka do obsługi smtp
	"net/smtp"
)

func main() {
	// Stworzenie klienta smtp
	c, err := smtp.Dial("127.0.0.1:10025")
	if err != nil {
		log.Fatal(err)
	}

	// Ustawieni pola Mail
	if err := c.Mail("user@devmeeting.org"); err != nil {
		log.Fatal(err)
	}

	// Ustawienie pola Rcpt
	if err := c.Rcpt("server@devmeeting.net"); err != nil {
		log.Fatal(err)
	}

	// Pobranie strumienia do zapisu body
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}

	// Wpisanie "hi" do body
	if _, err = fmt.Fprintf(wc, "hi"); err != nil {
		log.Fatal(err)
	}

	// Zamknięcie strumienia do zapisu
	if err = wc.Close(); err != nil {
		log.Fatal(err)
	}

	// Wysłanie Quit do sewera i tym samym wysłanie maila
	if err = c.Quit(); err != nil {
		log.Fatal(err)
	}
}
