package carbuilder

import "fmt"

type EngineType int

const (
	Manual EngineType = iota + 1
	Automatic
)

type Car struct {
	Model        string
	Color        string
	EngineType   EngineType
	NumberOfDoor int
}

func GetCarModel(model string) (IBuilder, error) {
	if model == "basic" {
		return &BasicCar{}, nil
	} else if model == "luxury" {
		return &LuxuryCar{}, nil
	}

	return nil, fmt.Errorf("car model not exist")
}

func GetCarDetail(model string, car Car) {
	fmt.Printf("----- %s -----\n", model)
	fmt.Printf("Model \t\t\t:%s\n", car.Model)
	fmt.Printf("Engine Type \t\t:%v\n", car.EngineType)
	fmt.Printf("Color \t\t\t:%s\n", car.Color)
	fmt.Printf("Number of Door \t\t:%d\n", car.NumberOfDoor)
}
