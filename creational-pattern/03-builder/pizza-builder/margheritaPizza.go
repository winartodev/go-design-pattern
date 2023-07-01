package pizzabuilder

type MargheritaPizzaBuilder struct {
	Size     PizzaSize
	Curst    CurstType
	Toppings ToppingType
	Sauce    SauceType
}

func (m *MargheritaPizzaBuilder) SetSize() {
	m.Size = Medium
}

func (m *MargheritaPizzaBuilder) SetCrustType() {
	m.Curst = ThinCrust
}

func (m *MargheritaPizzaBuilder) SetToppings() {
	m.Toppings = Seafood
}

func (m *MargheritaPizzaBuilder) SetSauce() {
	m.Sauce = WhiteSauce
}

func (m *MargheritaPizzaBuilder) Build() Pizza {
	return Pizza{
		Size:     m.Size,
		Curst:    m.Curst,
		Toppings: m.Toppings,
		Sauce:    m.Sauce,
	}
}
