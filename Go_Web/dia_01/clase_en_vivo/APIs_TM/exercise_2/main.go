package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

func main() {

	// Crea un router con gin
	router := gin.Default()

	// Captura la solicitud POST "/saludo"
	router.POST("/saludo", func(c *gin.Context) {
		var person Person
		if err := c.BindJSON(&person); err != nil {
			return
		}
		c.IndentedJSON(http.StatusOK, fmt.Sprintf("Hola %s %s", person.Nombre, person.Apellido))
	})

	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()
}
