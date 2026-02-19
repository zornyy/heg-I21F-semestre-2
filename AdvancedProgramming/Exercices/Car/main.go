package main

import (
	"car/car"
	"fmt"
)

func main() {
	var myCar car.Car
	myCar.SetModel("Super Modele de fdp")

	fmt.Printf("Car model: %s\n", myCar.GetModel())
}
