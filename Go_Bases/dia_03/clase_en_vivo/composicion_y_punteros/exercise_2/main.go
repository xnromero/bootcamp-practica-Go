package main

import "fmt"

type Producto interface {
	Precio() float64
}

type Pequenio struct{
	Price float64
}

type Mediano struct{
	Price float64
}

type Grande struct{
	Price float64
}

func main() {

	fmt.Println(factory("pequenio", 1000).Precio())
	fmt.Println(factory("mediano", 1000).Precio())
	fmt.Println(factory("grande", 1000).Precio())

}

func (p Pequenio) Precio() float64{
	return p.Price
}

func (m Mediano) Precio() float64{
	return m.Price + (m.Price * 0.03)
}

func (g Grande) Precio() float64{
	return g.Price + (g.Price * 0.06) + 2500
}


func factory(tipoProducto string, precio float64) Producto{
	switch tipoProducto{
	case "pequenio":
		return Pequenio{Price: precio}
	case "mediano":
		return Mediano{Price: precio}
	case "grande":
		return Grande{Price: precio}
	default:
		return nil
	}
}
