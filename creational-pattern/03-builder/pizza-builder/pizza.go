package pizzabuilder

import "fmt"

type PizzaSize int

const (
	Small PizzaSize = iota + 1
	Medium
	Large
)

func (c PizzaSize) String() string {
	return [...]string{"Small", "Medium", "Large"}[c-1]
}

type CurstType int

const (
	ThinCrust CurstType = iota + 1
	ThickCrust
	StuffedCrust
	GlutenFreeCrust
)

func (c CurstType) String() string {
	return [...]string{"Thin Crust", "Thick Crust", "Stuffed Crust", "Gluten Free Crust"}[c-1]
}

type ToppingType int

const (
	Cheese ToppingType = iota + 1
	Meats
	Vegetables
	Seafood
	Specialty
)

func (c ToppingType) String() string {
	return [...]string{"Cheese", "Meats", "Vegetables", "Seafood", "Specialty"}[c-1]
}

type SauceType int

const (
	WhiteSauce SauceType = iota + 1
	BarbecueSauce
	PestoSauce
)

func (c SauceType) String() string {
	return [...]string{"White Sauce", "Barbecue Sauce", "Pesto Sauce"}[c-1]
}

type Pizza struct {
	Size     PizzaSize
	Curst    CurstType
	Toppings ToppingType
	Sauce    SauceType
}

type PizzaType int

const (
	MargheritaPizza PizzaType = iota + 1
	PepperoniPizza
)

func (c PizzaType) String() string {
	return [...]string{"MargheritaPizza", "PepperoniPizza"}[c-1]
}

func GetPizzaType(pizzaType PizzaType) (PizzaBuilder, error) {
	if pizzaType == MargheritaPizza {
		return &MargheritaPizzaBuilder{}, nil
	} else if pizzaType == PepperoniPizza {
		return &PepperoniPizzaBuilder{}, nil
	}

	return nil, fmt.Errorf("pizza type not found")
}

func GetPizzaDetail(pizzaType string, pizza Pizza) {
	fmt.Printf("----- %s -----\n", pizzaType)
	fmt.Printf("Size \t\t:%s\n", pizza.Size.String())
	fmt.Printf("Curst \t\t:%s\n", pizza.Curst.String())
	fmt.Printf("Toppings \t:%s\n", pizza.Toppings.String())
	fmt.Printf("Sauce \t\t:%s\n", pizza.Sauce.String())
}
