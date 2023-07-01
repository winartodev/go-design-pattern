package pizzabuilder

type Director struct {
	builder PizzaBuilder
}

func NewPizzaDirector(b PizzaBuilder) Director {
	return Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b PizzaBuilder) {
	d.builder = b
}

func (d *Director) Build() Pizza {
	d.builder.SetSize()
	d.builder.SetCrustType()
	d.builder.SetSauce()
	d.builder.SetToppings()
	return d.builder.Build()
}
