package housebuilder

import "fmt"

type House struct {
	Windows     int
	Door        int
	Rooms       int
	HasGarage   bool
	HasSwimPool bool
	HasGarden   bool
}

func HouseDetail(houseType string, house House) {
	fmt.Printf("----- %s -----\n", houseType)
	fmt.Printf("Door \t\t\t:%d\n", house.Door)
	fmt.Printf("Windows \t\t:%d\n", house.Windows)
	fmt.Printf("Rooms \t\t\t:%d\n", house.Rooms)
	fmt.Printf("Has Garage \t\t:%v\n", house.HasGarage)
	fmt.Printf("Has Swimming Pool \t:%v\n", house.HasSwimPool)
	fmt.Printf("Has Garden \t\t:%v\n", house.HasGarden)
}
