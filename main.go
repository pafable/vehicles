package main

import (
	"cars/pkg/car"
	"fmt"
	"log"
)

func main() {
	beemer, err := car.GetCar("M340i")

	if err != nil {
		log.Println(err)
	}

	log.Println(beemer)
	beemer.SetDetail("speed", "fast")
	log.Println(beemer)

	fmt.Println(beemer.GetSpec("speed"))

	log.Println(beemer.GetMfr())
	log.Println(beemer.GetSpec("origin"))
}
