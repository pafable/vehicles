package car

import "fmt"

func GetCar(mfr string) (ICar, error) {
	switch mfr {
	case "bmw":
		return newBmw(), nil
	default:
		return nil, fmt.Errorf("no mfr entered")
	}
}
