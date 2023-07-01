package housebuilder

type IBuilder interface {
	SetWindows()
	SetDoor()
	SetRoom()
	SetGarage()
	SetSwimmingPool()
	SetGarden()
	GetHouse() House
}

func GetBuilderType(builderType string) IBuilder {
	if builderType == "normal" {
		return &NormalHouse{}
	} else if builderType == "iglo" {
		return &IgloHouse{}
	}

	return nil
}
