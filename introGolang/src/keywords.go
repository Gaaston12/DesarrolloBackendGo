package main

import "fmt"

func main() {
	// defer -> ejecuta la ultima accion antes de que el programa termine
	// por ejemplo al usar una bd, al abrir la coneccion luego se cierra con el defer.
	// usar 1 solo defer por funcion, como buena practica.
	defer fmt.Println("Hola")
	fmt.Println("mundo")

	// continue y br -> se usan adentro de un for

	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// continue -> se utiliza cuando una condificion dada dentro del ciclo te interesa que se continue
		// si sucede un error controlado
		if i == 2 {
			fmt.Println("es 2")
			continue
		}
		// break -> para cortar la ejecucion del programa en el momento donde se cumpla la condicion
		//
		if i == 8 {
			fmt.Println("es 8")
			break
		}
	}
}
