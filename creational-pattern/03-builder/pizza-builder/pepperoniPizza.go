package pizzabuilder

type PepperoniPizzaBuilder struct {
	Size     PizzaSize
	Curst    CurstType
	Toppings ToppingType
	Sauce    SauceType
}

func (p *PepperoniPizzaBuilder) SetSize() {
	p.Size = Large
}

func (p *PepperoniPizzaBuilder) SetCrustType() {
	p.Curst = StuffedCrust
}

func (p *PepperoniPizzaBuilder) SetToppings() {
	p.Toppings = Cheese
}

func (p *PepperoniPizzaBuilder) SetSauce() {
	p.Sauce = PestoSauce
}

func (p *PepperoniPizzaBuilder) Build() Pizza {
	return Pizza{
		Size:     p.Size,
		Curst:    p.Curst,
		Toppings: p.Toppings,
		Sauce:    p.Sauce,
	}
}
