package foodfactory

import "fmt"

type FoodItem interface {
	Prepare() string
}

type ItemType int

const (
	Pizza ItemType = iota + 1
	Burger
	Pasta
)

type Food struct {
	Name string
}

func (f *Food) Prepare() string {
	return fmt.Sprintf("Prepare %s", f.Name)
}

type PizzaItem struct {
	Food
}

type BurgerItem struct {
	Food
}

type PastaItem struct {
	Food
}

type FoodItemFactory struct{}

func (fif *FoodItemFactory) CreateFoodItem(itemType ItemType) (FoodItem, error) {
	if itemType == Pizza {
		return &PizzaItem{
			Food{
				Name: "Pizza",
			},
		}, nil
	} else if itemType == Burger {
		return &BurgerItem{
			Food{
				Name: "Burger",
			},
		}, nil
	} else if itemType == Pasta {
		return &PastaItem{
			Food{
				Name: "Pasta",
			},
		}, nil
	} else {
		return nil, fmt.Errorf("item type not found")
	}
}
