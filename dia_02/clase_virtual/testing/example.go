package main

import "fmt"

var users = []string {
	"Vincent",
	"Sol",
	"Nico",
	"Mati",
}

func main(){
	addUser("Agus")
	fmt.Println("Los usuarios registrados son: ", users)

}

func addUser(name string){
	users = append(users, name)
}