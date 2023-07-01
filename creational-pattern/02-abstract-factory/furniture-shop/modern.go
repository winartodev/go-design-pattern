package furnitureshop

type Modern struct {
}

func (ad *Modern) CreateChair() IChair {
	return &Chair{
		Color:        "yellow",
		NumberOfLegs: 4,
	}
}

func (ad Modern) CreateCoofeTable() ICoffeTable {
	return &CoffeTable{
		Length: 5,
		Width:  6,
	}
}

func (ad *Modern) CreateSofa() ISofa {
	return &Sofa{
		Color:    "green",
		Material: "hard",
	}
}
