package car

type Car struct {
	specs map[string]string
}

func (c *Car) SetDetail(key string, value string) {
	c.specs[key] = value
}

func (c *Car) GetMfr() string {
	return c.specs["mfr"]
}

func (c *Car) GetSpec(query string) string {
	return c.specs[query]
}
