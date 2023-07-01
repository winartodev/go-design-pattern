package carbuilder

type BasicCar struct {
	Model        string
	Color        string
	EngineType   EngineType
	NumberOfDoor int
}

func NewBasicCar() *BasicCar {
	return &BasicCar{}
}

func (bc *BasicCar) SetModel() {
	bc.Model = "Basic Car"
}

func (bc *BasicCar) SetColor() {
	bc.Color = "Black"
}

func (bc *BasicCar) SetEngine() {
	bc.EngineType = Manual
}

func (bc *BasicCar) SetNumberOfDoor() {
	bc.NumberOfDoor = 4
}

func (bc *BasicCar) Build() Car {
	return Car{
		Model:        bc.Model,
		Color:        bc.Color,
		EngineType:   bc.EngineType,
		NumberOfDoor: bc.NumberOfDoor,
	}
}
