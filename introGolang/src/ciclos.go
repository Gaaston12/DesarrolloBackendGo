package main

import "fmt"

func main() {

	// for condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for while
	fmt.Println()
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// for forever
	// fmt.Println()
	// counterForever := 0
	// for {
	// 	fmt.Println(counterForever)
	// 	counterForever++
	// }

	//ejercicio
	counter2 := 10
	for counter2 >= 0 {
		fmt.Println(counter2)
		counter2--
	}

}
