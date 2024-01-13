package car

import "fmt"

func GetCar(model string) (ICar, error) {
	switch model {
	case "M340i":
		return newBmw(), nil
	default:
		return nil, fmt.Errorf("no mfr entered")
	}
}
