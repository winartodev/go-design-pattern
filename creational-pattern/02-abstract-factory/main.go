package main

import (
	"fmt"

	characterabstractfactory "github.com/winartodev/go-design-pattern/creational-pattern/02-abstract-factory/character-abstract-factory"
	furnitureshop "github.com/winartodev/go-design-pattern/creational-pattern/02-abstract-factory/furniture-shop"
)

func main() {
	warior, _ := characterabstractfactory.GetCharacterFactory(characterabstractfactory.Warrior)
	resultWarior := warior.MakeCharacter()
	fmt.Println(resultWarior.GetCharacterType())
	fmt.Println(resultWarior.PerfomAction())

	fmt.Println()

	evil, _ := characterabstractfactory.GetCharacterFactory(characterabstractfactory.Mage)
	resultEvil := evil.MakeCharacter()
	fmt.Println(resultEvil.GetCharacterType())
	fmt.Println(resultEvil.PerfomAction())

	artDeco, _ := furnitureshop.GetFurnitureVariant(furnitureshop.ArtDecoType)
	artDecoChair := artDeco.CreateChair()
	artDecoCoffeTable := artDeco.CreateCoofeTable()
	artDecoSofa := artDeco.CreateSofa()
	furnitureShopDetail(furnitureshop.ArtDecoType.String(), artDecoChair, artDecoCoffeTable, artDecoSofa)

	victorian, _ := furnitureshop.GetFurnitureVariant(furnitureshop.VictorianType)
	victorianChair := victorian.CreateChair()
	victorianCoffeTable := victorian.CreateCoofeTable()
	victorianSofa := victorian.CreateSofa()
	furnitureShopDetail(furnitureshop.VictorianType.String(), victorianChair, victorianCoffeTable, victorianSofa)

	modern, _ := furnitureshop.GetFurnitureVariant(furnitureshop.ModernType)
	modernChair := modern.CreateChair()
	modernCoffeTable := modern.CreateCoofeTable()
	modernSofa := modern.CreateSofa()
	furnitureShopDetail(furnitureshop.ModernType.String(), modernChair, modernCoffeTable, modernSofa)
}

func furnitureShopDetail(varianType string, chair furnitureshop.IChair, coffeTable furnitureshop.ICoffeTable, sofa furnitureshop.ISofa) {
	fmt.Println()
	fmt.Printf("----- %s Chair -----\n", varianType)
	fmt.Printf("%s Chair Color\t\t: %s\n", varianType, chair.GetColor())
	fmt.Printf("%s Chair Num Of Legs\t: %d\n", varianType, chair.GetNumberOfLegs())
	fmt.Println()

	fmt.Printf("----- %s Coffe Table -----\n", varianType)
	fmt.Printf("%s Coffe Table Width\t: %d\n", varianType, coffeTable.GetWidth())
	fmt.Printf("%s Coffe Table Length\t: %d\n", varianType, coffeTable.GetLength())
	fmt.Println()

	fmt.Printf("----- %s Sofa -----\n", varianType)
	fmt.Printf("%s Sofa Color\t\t: %s\n", varianType, sofa.GetColor())
	fmt.Printf("%s Sofa Material\t\t: %s\n", varianType, sofa.GetMaterial())
	fmt.Println()
}
