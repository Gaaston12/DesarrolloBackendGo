package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["gas"] = 23
	m["martin"] = 22

	fmt.Println(m)

	//recorer un map
	for _, v := range m {
		fmt.Println(v)
	}

	//encontrar un valor
	value, ok := m["gas"]
	if !ok {
		fmt.Println("no existe")
	}
	fmt.Println(value)
}
