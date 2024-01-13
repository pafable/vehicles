package car

type Bmw struct {
	Car
}

func newBmw() ICar {
	return &Bmw{
		Car: Car{
			specs: map[string]string{
				"mfr":    "bmw",
				"origin": "DE",
			},
		},
	}
}
