package car

type Car struct {
	model string
}

func (c *Car) SetModel(model string) {
	c.model = model
}

func (c Car) GetModel() string {
	return c.model
}
