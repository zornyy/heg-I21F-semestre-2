package main

import (
	"bicycle/bicycle"
	"fmt"
)

func main() {
	var b bicycle.Bicycle

	b.SetSpeed(5)
	b.SetGear(2)

	fmt.Printf("The bike is moving at speed %d in gear %d\n", b.GetSpeed(), b.GetGear())
}
