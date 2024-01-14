package plane

import "fmt"

func GetPlane(model string, mach float32) (IPlane, error) {
	switch model {
	case "F22":
		return newLockheed(mach), nil
	default:
		return nil, fmt.Errorf("no mfr found")
	}
}
