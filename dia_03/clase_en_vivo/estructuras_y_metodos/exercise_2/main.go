package main

import (
	//"encoding/json"
	"fmt"
)

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	position string
	Person
}

func main() {

	person := Person{1212, "Ximena", "14-12-1993"}
	employee := Employee{1534, "Desarrollador de Software", person}

	employee.PrintEmployee()

}

func (e Employee) PrintEmployee() {
	//JSON, _ := (json.Marshal(e))
	//fmt.Println(string(JSON))
	fmt.Println("Id Person: ", e.Person.ID, ", Name: ", e.Name, ", Date of birth: ", e.DateOfBirth,
		", Id Employee: ", e.ID, ", Position: ", e.position)
}
