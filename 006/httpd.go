// nazwa paczki
package main

// list importowanych paczek
import (
	// standardowe paczki go
	"log"

	// paczk serwera http
	// pobrać zewnętrzną paczkę można za pomocą komendy:
	// go get github.com/gin-gonic/gin
	"github.com/gin-gonic/gin"
)

// funkcja main
func main() {
	// deklaracja routera
	router := gin.Default()

	// metoda GET
	router.GET("/", func(c *gin.Context) {
		c.String(200, "GET replay")
	})

	// metoda POST
	// :email - tak deklarujemy parametry w path
	router.POST("/insert/:email", func(c *gin.Context) {
		c.String(200, "POST "+c.Param("email")+" replay")
	})

	// metoda DELETE
	router.DELETE("/remove", func(c *gin.Context) {
		c.String(200, "DELETE replay")
	})

	// Uruchomienie serwera na porcie 8080 oraz logowanie ewentualnych błędów
	log.Fatal(router.Run(":8080"))
}
