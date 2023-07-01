package housebuilder

type NormalHouse struct {
	Windows     int
	Door        int
	Rooms       int
	HasGarage   bool
	HasSwimPool bool
	HasGarden   bool
}

func NewNormalHouse() *NormalHouse {
	return &NormalHouse{}
}

func (nb *NormalHouse) SetWindows() {
	nb.Windows = 6
}

func (nb *NormalHouse) SetDoor() {
	nb.Door = 3
}

func (nb *NormalHouse) SetRoom() {
	nb.Rooms = 3
}

func (nb *NormalHouse) SetGarage() {
	nb.HasGarage = true
}

func (nb *NormalHouse) SetSwimmingPool() {
	nb.HasSwimPool = false
}

func (nb *NormalHouse) SetGarden() {
	nb.HasGarden = false
}

func (nb *NormalHouse) GetHouse() House {
	return House{
		Windows:     nb.Windows,
		Door:        nb.Door,
		Rooms:       nb.Rooms,
		HasGarage:   nb.HasGarage,
		HasSwimPool: nb.HasSwimPool,
		HasGarden:   nb.HasGarden,
	}
}
