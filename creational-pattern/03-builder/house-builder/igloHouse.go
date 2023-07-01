package housebuilder

type IgloHouse struct {
	Windows     int
	Door        int
	Rooms       int
	HasGarage   bool
	HasSwimPool bool
	HasGarden   bool
}

func NewIgloHouse() {

}

func (ih *IgloHouse) Reset() {

}

func (ih *IgloHouse) SetWindows() {
	ih.Windows = 2
}

func (ih *IgloHouse) SetDoor() {
	ih.Door = 1
}

func (ih *IgloHouse) SetRoom() {
	ih.Rooms = 1
}

func (ih *IgloHouse) SetGarage() {
	ih.HasGarage = false
}

func (ih *IgloHouse) SetSwimmingPool() {
	ih.HasSwimPool = false
}

func (ih *IgloHouse) SetGarden() {
	ih.HasGarden = false
}

func (ih *IgloHouse) GetHouse() House {
	return House{
		Windows:     ih.Windows,
		Door:        ih.Door,
		Rooms:       ih.Rooms,
		HasGarage:   ih.HasGarage,
		HasSwimPool: ih.HasSwimPool,
		HasGarden:   ih.HasGarden,
	}
}
