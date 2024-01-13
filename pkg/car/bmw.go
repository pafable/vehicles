package car

type Bmw struct {
	Car
}

func newBmw() ICar {
	return &Bmw{
		Car{
			model: "M340i",
		},
	}
}
