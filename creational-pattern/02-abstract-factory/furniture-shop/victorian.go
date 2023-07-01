package furnitureshop

type Victorian struct {
}

func (v *Victorian) CreateChair() IChair {
	return &Chair{
		Color:        "Red",
		NumberOfLegs: 4,
	}
}

func (v Victorian) CreateCoofeTable() ICoffeTable {
	return &CoffeTable{
		Length: 2,
		Width:  3,
	}
}

func (v *Victorian) CreateSofa() ISofa {
	return &Sofa{
		Color:    "Blue",
		Material: "Soft",
	}
}
