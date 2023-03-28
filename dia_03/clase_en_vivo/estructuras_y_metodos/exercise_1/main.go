package main

import (
	"fmt"
)

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products = []Product{}

func main() {

	product1 := Product{1231, "Coca Cola", 19.90, "Bebida gaseosa", " "}
	product2 := Product{1232, "Agua", 11.00, "Agua Mineral", " "}

	fmt.Println(Products)
	fmt.Println("Agrego producto")
	product1.Save()
	GetAll()
	fmt.Println("Agrego producto")
	product2.Save()
	GetAll()
	fmt.Println(getById(1231))


}

func (p Product) Save() {

	Products = append(Products, p)

}

func GetAll() {

	for _, product := range Products{
		fmt.Println(product)
	}

}

func getById(id int) Product{
    var product Product
	for i := range Products{
		if Products[i].ID == id {
			product = Products[i]
			break
		}
	}
	return product
}
