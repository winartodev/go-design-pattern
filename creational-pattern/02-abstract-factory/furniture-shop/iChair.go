package furnitureshop

type IChair interface {
	GetColor() string
	GetNumberOfLegs() int
}

type Chair struct {
	Color        string
	NumberOfLegs int
}

func (c *Chair) GetColor() string {
	return c.Color
}

func (c *Chair) GetNumberOfLegs() int {
	return c.NumberOfLegs
}
