package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// Crea un router con gin
	router := gin.Default()

	// Captura la solicitud GET "/ping"
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()
}
