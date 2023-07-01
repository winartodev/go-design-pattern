package main

import (
	"fmt"

	bookbuilder "github.com/winartodev/go-design-pattern/creational-pattern/03-builder/book-builder"
	carbuilder "github.com/winartodev/go-design-pattern/creational-pattern/03-builder/car-builder"
	housebuilder "github.com/winartodev/go-design-pattern/creational-pattern/03-builder/house-builder"
	pizzabuilder "github.com/winartodev/go-design-pattern/creational-pattern/03-builder/pizza-builder"
)

func main() {
	makeNormalHouse := housebuilder.GetBuilderType("normal")
	makeIgloHouse := housebuilder.GetBuilderType("iglo")

	director := housebuilder.NewDirector(makeNormalHouse)
	normalHouse := director.BuildHouse()
	housebuilder.HouseDetail("Normal House", normalHouse)

	fmt.Println()
	director.SetBuilder(makeIgloHouse)
	igloHouse := director.BuildHouse()
	housebuilder.HouseDetail("Iglo House", igloHouse)

	fmt.Println()
	basicCarModel, _ := carbuilder.GetCarModel("basic")
	luxuryCarModel, _ := carbuilder.GetCarModel("luxury")

	carDirector := carbuilder.NewDirector(basicCarModel)
	basicCar := carDirector.Build()
	carbuilder.GetCarDetail("Basic Car", basicCar)

	fmt.Println()
	carDirector.SetBuilder(luxuryCarModel)
	luxuryCar := carDirector.Build()
	carbuilder.GetCarDetail("Luxury Car", luxuryCar)

	makePizza1, _ := pizzabuilder.GetPizzaType(pizzabuilder.MargheritaPizza)
	makePizza2, _ := pizzabuilder.GetPizzaType(pizzabuilder.PepperoniPizza)

	pizzaDirector := pizzabuilder.NewPizzaDirector(makePizza1)
	margheritaPizza := pizzaDirector.Build()

	pizzaDirector.SetBuilder(makePizza2)
	paperoniPizza := pizzaDirector.Build()

	fmt.Println()
	pizzabuilder.GetPizzaDetail(pizzabuilder.MargheritaPizza.String(), margheritaPizza)
	fmt.Println()
	pizzabuilder.GetPizzaDetail(pizzabuilder.PepperoniPizza.String(), paperoniPizza)

	bookBuilder := &bookbuilder.BookBuilder{}
	book1 := bookBuilder.
		SetTitle("The Alchemist").
		SetAuthor("Paulo Coelho").
		SetISBN("9780062315007").
		Build()

	book2 := bookBuilder.
		SetTitle("To Kill a Mockingbird").
		SetAuthor("Harper Lee").
		SetISBN("978-0061120084").
		SetPublicationYear(1960).
		SetPublisher("Harper Perennial Modern Classics").
		SetDescription(`"To Kill a Mockingbird" is a renowned coming-of-age novel set in the racially divided Deep South during the 1930s. It follows the story of Scout Finch, a young girl, as she learns about compassion, empathy, and the injustices of her society.`).
		Build()

	book1.PrintInfo()
	book2.PrintInfo()
}
