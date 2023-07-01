package housebuilder

type Director struct {
	builder IBuilder
}

func NewDirector(builder IBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) SetBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) BuildHouse() House {
	d.builder.SetWindows()
	d.builder.SetDoor()
	d.builder.SetRoom()
	d.builder.SetGarage()
	d.builder.SetSwimmingPool()
	d.builder.SetGarden()

	return d.builder.GetHouse()
}
