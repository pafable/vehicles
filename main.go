package main

import (
	"github.com/pafable/cars/pkg/car"
	"github.com/pafable/cars/pkg/plane"
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

	log.Println(beemer.GetSpec("speed"))

	log.Println(beemer.GetMfr())
	log.Println(beemer.GetSpec("origin"))

	raptor, err := plane.GetPlane("F22", 2.2)
	raptor.SetDetail("pilot", "1")
	log.Println(raptor)
	log.Println(raptor.GetMach())
	log.Println(raptor.GetSpec("pilot"))
}
