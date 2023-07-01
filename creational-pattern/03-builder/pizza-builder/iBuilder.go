package pizzabuilder

type PizzaBuilder interface {
	SetSize()
	SetCrustType()
	SetToppings()
	SetSauce()
	Build() Pizza
}
