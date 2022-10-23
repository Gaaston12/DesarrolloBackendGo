package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	fmt.Println(myPc.brand, "Pong")
}
func (myPc *pc) duplicateRam() {
	myPc.ram = myPc.ram * 2
}

func main() {
	a := 10
	b := &a
	fmt.Println(a)
	fmt.Println(*b)

	myPc := pc{ram: 16, disk: 5, brand: "mac"}
	fmt.Println(myPc)

	//myPc.ping()
	myPc.duplicateRam()
	fmt.Println(myPc)
	myPc.duplicateRam()
	fmt.Println(myPc)
	fmt.Printf("2, %T", 2)
}
