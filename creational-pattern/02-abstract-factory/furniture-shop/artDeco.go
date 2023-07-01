package furnitureshop

type ArtDeco struct {
}

func (ad *ArtDeco) CreateChair() IChair {
	return &Chair{
		Color:        "green",
		NumberOfLegs: 2,
	}
}

func (ad ArtDeco) CreateCoofeTable() ICoffeTable {
	return &CoffeTable{
		Length: 3,
		Width:  4,
	}
}

func (ad *ArtDeco) CreateSofa() ISofa {
	return &Sofa{
		Color:    "white",
		Material: "Medium",
	}
}
