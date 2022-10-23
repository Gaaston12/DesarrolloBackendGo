package main

import (
	"fmt"
	"strings"
)

func main() {

	// array -> INMUTABLE!!
	var arr [4]int
	arr[0] = 1
	fmt.Println(arr)
	fmt.Printf("cantidad de elementos: %d", len(arr))
	fmt.Printf("capacidad maxima del array: %d", cap(arr))

	//slice -> MUTABLE!!
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// metodos en slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])  // el 3 no se muestra, porque el ultimo elemento no esta incluido
	fmt.Println(slice[2:4]) // estos 2 si son incluidos
	fmt.Println(slice[4:])  // el primer elmento esta incluido

	// append
	slice = append(slice, 7)
	fmt.Println(slice)

	// append new list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	// recorer un slice
	sliceS := []string{"q", "w", "e", "r", "t", "y"}
	for _, valor := range sliceS {
		fmt.Println(valor)
	}

	// palindomo
	// amor a roma - ama
	isPalindromo(strings.ToLower("Am0r a r0ma"))

	//años := make(map[int][]string)
}

func isPalindromo(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("es palindromo")
	} else {
		fmt.Println("no es palindromo")
	}
}
