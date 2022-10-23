package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("es 1")
	} else {
		fmt.Println("no es 1")
	}

	// With and
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("es verdad")
	}

	// With or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("al menos 1 es verdad")
	}

	// convertir texto a num
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)

	// reto 1
	num := 2
	if num%2 == 0 {
		fmt.Printf("el numero %d es par\n", num)
	} else {
		fmt.Printf("el numero %d es impar\n", num)
	}

	// reto 2

	user := "user"
	pass := "password"

	if user != "user" || pass != "password" {
		fmt.Println("usuario o contraseña incorrecta")
	}

	// switch

	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("es par")
	default:
		fmt.Println("es impar")
	}

	
}
