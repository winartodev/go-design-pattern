package carbuilder

type LuxuryCar struct {
	Model        string
	Color        string
	EngineType   EngineType
	NumberOfDoor int
}

func NewLuxuryCar() *LuxuryCar {
	return &LuxuryCar{}
}

func (lc *LuxuryCar) SetModel() {
	lc.Model = "Luxury Car"
}

func (lc *LuxuryCar) SetColor() {
	lc.Color = "White"
}

func (lc *LuxuryCar) SetEngine() {
	lc.EngineType = Automatic
}

func (lc *LuxuryCar) SetNumberOfDoor() {
	lc.NumberOfDoor = 2
}

func (lc *LuxuryCar) Build() Car {
	return Car{
		Model:        lc.Model,
		Color:        lc.Color,
		EngineType:   lc.EngineType,
		NumberOfDoor: lc.NumberOfDoor,
	}
}
