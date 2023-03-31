package main

import "fmt"

func main() {
	//Counter()
	//IterateArray()
	//IterateSlice()
	//IterateMap()
	Iterate()
}

//Incrementador de a 2 que solo se repita 5 veces
func Counter() {
	var lim = 5
	num := 10

	for i := 0; i < lim; i++ {
		num += 2
	}
	fmt.Println(num)
}

//Crear array de items, mostrar el primero y el último
func IterateArray() {
	frutas := [4] string {"banana", "manzana", "pera", "uva"}
	fmt.Println("len: ", len(frutas), " cap: ", cap(frutas))

	fmt.Println("Primer item: ", frutas[0])
	fmt.Println("Último item: ", frutas[len(frutas)-1])

}

//Crear un slice, mostrar el primer y último item.
func IterateSlice(){
	frutas:= []string {"banana", "manzana", "pera", "uva"}

	fmt.Println("len: ", len(frutas), " cap: ", cap(frutas))
	fmt.Println("Primer item: ", frutas[0])
	fmt.Println("Último item: ", frutas[len(frutas)-1])

	//agregar in item
	frutas = append(frutas, "naranja")
	fmt.Println("len: ", len(frutas), " cap: ", cap(frutas))

	//iterar
	fmt.Println("Todos los items: ")
	for ix, item := range frutas {
		fmt.Println(ix, item)
	}

	// nuevo slice con make
	fmt.Println("\nNuevo Slice")
	sl := make([]int, 0, 5)       // len: 0, cap: 5
	fmt.Println("sl: ", sl)
	fmt.Println("len: ", len(sl), "cap: ", cap(sl))

	sl = append(sl, 1, 2, 3, 4, 5, 6)
	fmt.Println("sl: ", sl)
	fmt.Println("len: ", len(sl), "cap: ", cap(sl))

}

// Crear un registro de persona y sus items
func IterateMap() {
	//crear map, agregar persona e imprimir
	var registro = make(map[string] []string)
	registro["Juan"] = []string{"banana", "manzana", "pera", "uva"}
	fmt.Println("Registro: ", registro)

	//Validar si el registro existe y almacenar sus item para imprimirlos
	items, ok := registro["Juan"]
	if ok {
		fmt.Println("Items de Juan: ", items)
	}

	//Borrar persona
	fmt.Println("Se elimina Juan")
	delete(registro, "Juan")
	fmt.Println(registro)

	//Iterar
	fmt.Println("Agregando mas personas")
	registro["pedro"] = []string{"banana", "manzana", "pera", "uva"}
	registro["maria"] = []string{"naraja", "limon", "sandia", "melon"}

	for persona, items := range registro {
		fmt.Println("Persona: ", persona)
		fmt.Println("Items: ", items)
	}

}
//contar cuantas veces se repite un item en un slice
func Iterate() {
	var frutas = []string {"banana", "manzana", "pera", "uva", "banana", "manzana", "pera", "uva", "banana", "manzana", "melon"}
	var hashMap = make(map[string] int)

	for _, item := range frutas {
		hashMap[item]++
	}

	fmt.Println("HashMap: ", hashMap)
}