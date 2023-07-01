package furnitureshop

type ISofa interface {
	GetColor() string
	GetMaterial() string
}

type Sofa struct {
	Color    string
	Material string
}

func (s *Sofa) GetColor() string {
	return s.Color
}

func (s *Sofa) GetMaterial() string {
	return s.Material
}
