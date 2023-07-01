package main

import (
	"fmt"

	carfactory "github.com/winartodev/go-design-pattern/creational-pattern/01-factory-method/car-factory"
)

func main() {
	carFactory := carfactory.CarFactory{}

	SedanModel, _ := carFactory.CreateCar(carfactory.Sedan)
	fmt.Println(SedanModel.Produce())

	SUVModel, _ := carFactory.CreateCar(carfactory.SUV)
	fmt.Println(SUVModel.Produce())

	HatchbackModel, _ := carFactory.CreateCar(carfactory.Hatchback)
	fmt.Println(HatchbackModel.Produce())
}
