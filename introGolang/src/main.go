package main

import (
	"fmt"
	"math"
)

func main() {
	// hellw world
	fmt.Println("Hello world!!")

	// declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1416
	fmt.Println(pi, pi2)

	//decalracion de variables
	base := 12          // declarando la variable y asignando el valor
	var altura int = 14 // declarando la variable con el tipo e inicializandola
	var area int        // declarando la variable con el tipo pero sin inicializarla

	fmt.Println(base, altura, area)

	// ZERO values
	var a int     // 0
	var b float64 // 0
	var c string  // " "
	var d bool    // false

	fmt.Println(a, b, c, d)

	// area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("El area del cuadrado es: ", areaCuadrado)

	// operadores aritmeticos
	x := 10
	y := 50

	//suma
	result := x + y
	fmt.Println("Suma: ", result)

	//resta
	result = y - x
	fmt.Println("Resta: ", result)

	//multiplicacion
	result = y * x
	fmt.Println("Multiplicacion:", result)

	//divicion
	result = y / x
	fmt.Println("Divicion: ", result)

	//modulo
	result = y % x
	fmt.Println("Modulo: ", result)

	// calcular areas de
	// 		rectangulo trapecio circulo

	//rectangulo
	// const largo = 10
	// const ancho = 20

	// areaRectangulo := largo * ancho
	// fmt.Println("El area del rectangulo es: ", areaRectangulo)
	areaR := areaRectangulo(10, 20)
	fmt.Println("El area del rectangulo es: ", areaR)

	//trapecio
	// const a1 = 10
	// const b1 = 20
	// const altura1 = 30

	// areaTrapecio := ((a1 + b1) / 2) * altura1
	// fmt.Println("El area del rectangulo es: ", areaTrapecio)
	areaT := areaTrapecio(10, 20, 30)
	fmt.Println("El area del trapecio es: ", areaT)

	//circulo
	// const radio = 12.3

	// areaCirculo := (pi * math.Pow(float64(radio), 2))
	// fmt.Println("Area del circulo: ", areaCirculo)
	areaC := areaCirculo(12.3)
	fmt.Println("Area del circulo: ", areaC)

	// GO DOC
	// https://pkg.go.dev/std
	// https://pkg.go.dev/
}

func areaRectangulo(alto, ancho float64) float64 {
	return alto * ancho
}

func areaTrapecio(a, b, altura float64) float64 {
	return ((a + b) / 2) * altura
}

func areaCirculo(radio float64) float64 {
	const pi = 3.14
	return (pi * math.Pow(float64(radio), 2))
}
