package carfactory

import "fmt"

type ICar interface {
	Produce() string
}

type CarTypeEnum int

const (
	Sedan CarTypeEnum = iota + 1
	SUV
	Hatchback
)

func (c CarTypeEnum) String() string {
	return [...]string{"Sedan", "SUV", "Hatchback"}[c-1]
}

type Car struct {
	CarTypeEnum CarTypeEnum
}

func (c *Car) Produce() string {
	return fmt.Sprintf("Produce %s car", c.CarTypeEnum.String())
}

type SedanModel struct {
	Car
}

type SUVModel struct {
	Car
}

type HatchbackModel struct {
	Car
}

type CarFactory struct{}

func (c *CarFactory) CreateCar(carType CarTypeEnum) (ICar, error) {
	switch carType {
	case Sedan:
		return &SedanModel{
			Car{
				CarTypeEnum: Sedan,
			},
		}, nil
	case SUV:
		return &SUVModel{
			Car{
				CarTypeEnum: SUV,
			},
		}, nil
	case Hatchback:
		return &HatchbackModel{
			Car{
				CarTypeEnum: Hatchback,
			},
		}, nil
	}

	return nil, fmt.Errorf("car type not found")
}
