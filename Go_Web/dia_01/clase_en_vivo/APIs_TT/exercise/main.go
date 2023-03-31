package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var Productos = []Producto{}

func main() {

	//Leer desde archivo json y cargar en Slice
	file, err := ioutil.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &Productos)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(Productos)

	router := gin.Default()

	router.GET("/ping", Pong)

	group := router.Group("/products")
	group.GET("/", Products)
	group.GET("/:id", ProductById)
	group.GET("/search", Search)

	router.Run()
}

func Pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func Products(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Productos)
}

func ProductById(c *gin.Context) {

	var Id int

	if value, err := strconv.Atoi(c.Param("id")); err == nil {
		Id = value
	}

	producto, ok := ProductExist(Id)
	if ok {
		c.IndentedJSON(http.StatusOK, producto)
	} else {
		c.IndentedJSON(http.StatusNotFound, nil)
	}
}

func ProductExist(index int) (Producto, bool) {

	var producto Producto

	for i := range Productos {
		if Productos[i].Id == index {
			return Productos[i], true
		}
	}
	return producto , false
}

func Search(c *gin.Context) {

	price, err := strconv.ParseFloat(c.Query("priceGt"), 64)
	if err != nil {
		log.Fatal(err)
	}
	 
	c.IndentedJSON(http.StatusOK, FindProducts(price))

}

func FindProducts(price float64) []Producto{

	var products []Producto

	for i := range Productos {
		if Productos[i].Price > price {
			products = append(products, Productos[i])
		}
	}

	return products
}
