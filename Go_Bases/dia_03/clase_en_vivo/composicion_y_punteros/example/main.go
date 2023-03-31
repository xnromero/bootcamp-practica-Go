package main

import (
	"fmt"
)

func main (){
	r := rect{3, 4}
	details(r)
}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct{
	width, height float64
}

func (r rect) area() float64{
	return r.width * r.height
}

func (r rect) perim() float64{
	return (2*r.width) + (2*r.height)
}

func details (g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

