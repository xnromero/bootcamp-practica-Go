package main

import (
	//"encoding/json"
	//"fmt"
	"github.com/gin-gonic/gin"
)


func main() {

	/*data := map[string]any{
		"name":      "John",
		"edad":      30,
		"isMarried": true,
		"shopList": []string{
			"milk",
			"apple",
			"coffee",
		},
	}

	
	dataAsJson, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Data as map: %+v", data)
	fmt.Printf("Data as Json: %s", dataAsJson)*/



		// Crea un router con gin
		router := gin.Default()

		// Captura la solicitud GET "/hello-world"
		router.GET("/hello-world", func(c *gin.Context){
			c.JSON(200, gin.H{
				"message": "Hello World",
			})
		})

		// Corremos nuestro servidor sobre el puerto 8080
		router.Run()
	
}
