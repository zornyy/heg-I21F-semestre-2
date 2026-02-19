package car

type Car struct {
	Model string
}

func (c *Car) SetModel(model string) {
	c.Model = model
}

func (c Car) GetModel() string {
	return c.Model
}
