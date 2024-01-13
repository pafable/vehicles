package car

type Car struct {
	model string
}

func (c Car) setModel(model string) {
	c.model = model
}

func (c Car) getModel() string {
	return c.model
}
