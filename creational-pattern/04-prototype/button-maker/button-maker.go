package buttonmaker

type IButtonMaker interface {
	Clone() IButtonMaker
	GetColor() string
}

type Button struct {
	X     float64
	Y     float64
	Color string
}

func (b *Button) Clone() IButtonMaker {
	return &Button{
		X:     b.X,
		Y:     b.Y,
		Color: b.Color,
	}
}

func (b *Button) GetColor() string {
	return b.Color
}

type ButtonRegistry struct {
	prototypes map[string]IButtonMaker
}

func NewButtonRegistry() *ButtonRegistry {
	return &ButtonRegistry{
		prototypes: make(map[string]IButtonMaker),
	}
}

func (b *ButtonRegistry) AddItem(id string, p IButtonMaker) {
	b.prototypes[id] = p
}

func (b *ButtonRegistry) GetByID(id string) IButtonMaker {
	return b.prototypes[id]
}

func (b *ButtonRegistry) GetByColor(color string) IButtonMaker {
	for _, item := range b.prototypes {
		if item.GetColor() == color {
			return item.Clone()
		}
	}

	return nil
}
