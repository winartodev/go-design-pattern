package carbuilder

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

func (d *Director) Build() Car {
	d.builder.SetModel()
	d.builder.SetColor()
	d.builder.SetEngine()
	d.builder.SetNumberOfDoor()
	return d.builder.Build()
}
