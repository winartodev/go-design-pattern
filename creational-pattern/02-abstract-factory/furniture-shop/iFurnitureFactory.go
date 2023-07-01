package furnitureshop

import "fmt"

type FurnitureVarianType int

const (
	ArtDecoType FurnitureVarianType = iota + 1
	VictorianType
	ModernType
)

func (c FurnitureVarianType) String() string {
	return [...]string{"Art Deco", "Victorian", "Modern"}[c-1]
}

type IFurnitureFactory interface {
	CreateChair() IChair
	CreateCoofeTable() ICoffeTable
	CreateSofa() ISofa
}

func GetFurnitureVariant(variant FurnitureVarianType) (IFurnitureFactory, error) {
	if variant == ArtDecoType {
		return &ArtDeco{}, nil
	} else if variant == VictorianType {
		return &Victorian{}, nil
	} else if variant == ModernType {
		return &Modern{}, nil
	}

	return nil, fmt.Errorf("variant type not exist")
}
