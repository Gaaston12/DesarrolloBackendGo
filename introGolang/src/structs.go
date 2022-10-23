package main

import "fmt"

// si la primer letra es mayuscula es publico... Se puede acceder desde cualquier paquete
// si es minuscula, es privado
type car struct { // -> equivalente a las clases de otros lenguajes
	brand string
	year  int
}

// si es publico, es buena practica colocar el nombre en un comentario y explicar que hace
// Dog dog con acceso pubico
type Dog struct {
	Name string
}

func main() {
	myCar := car{brand: "ford", year: 2020}
	fmt.Println(myCar)

	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
