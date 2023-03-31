package main

import (
	"fmt"
)

type Alumno struct{
	Nombre, Apellido, DNI, Fecha string 
}

func main(){

	alumno1 := newAlumno("Mateo", "Perez", "40.384.585", "06-07-2022")
	alumno2 := newAlumno("Mariana", "Gomez", "45.383.374", "06-04-2022")
	alumno1.Detalle()
	alumno2.Detalle()

}

func newAlumno (nombre, apellido, dni, fecha string) Alumno{
	return Alumno{
		Nombre: nombre,
		Apellido: apellido,
		DNI: dni,
		Fecha: fecha,
	}
}

func (a *Alumno) Detalle(){
	fmt.Printf("\nNombre:  %s \nApellido: %s \nDNI:  %s \nFecha: %s \n", a.Nombre, a.Apellido, a.DNI, a.Fecha ) 
}