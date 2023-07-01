package furnitureshop

type ICoffeTable interface {
	GetLength() int
	GetWidth() int
}

type CoffeTable struct {
	Length int
	Width  int
}

func (ct *CoffeTable) GetLength() int {
	return ct.Length
}

func (ct *CoffeTable) GetWidth() int {
	return ct.Width
}
