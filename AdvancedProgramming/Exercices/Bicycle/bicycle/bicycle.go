package bicycle

type Bicycle struct {
	speed int
	gear  int
}

func (b Bicycle) GetSpeed() int {
	return b.speed
}

func (b *Bicycle) SetSpeed(speed int) {
	b.speed = speed
}

func (b Bicycle) GetGear() int {
	return b.gear
}

func (b *Bicycle) SetGear(gear int) {
	b.gear = gear
}
