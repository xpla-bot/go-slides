package main

import (
	"log"

	// importowanie implementacji serwera smtpd
	"github.com/krhubert/smtpd"
)

func main() {
	// deklaracja serwera smtpd
	server := smtpd.Server{
		// Ustawienie handlera dla otrzymanych wiadomości
		// Handler musi przyjmować dwa argument - Peer i Envelope
		// oraz zwracać error
		Handler: func(peer smtpd.Peer, env smtpd.Envelope) error {
			// logowanie nadawcy
			log.Printf("got email from %s", env.Sender)
			return nil
		},
	}
	// wystartowanie serwera smtpd
	server.ListenAndServe("127.0.0.1:10025")
}
