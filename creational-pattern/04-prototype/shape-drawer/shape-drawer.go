package shapedrawer

import "fmt"

type Shape interface {
	Clone() Shape
	Draw()
}

type Circle struct {
	Radius float64
}

func (c *Circle) Clone() Shape {
	return &Circle{
		Radius: c.Radius,
	}
}

func (c *Circle) Draw() {
	fmt.Printf("Drawing a circle with radius: %v\n", c.Radius)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Clone() Shape {
	return &Rectangle{
		Width:  r.Width,
		Height: r.Height,
	}
}

func (r *Rectangle) Draw() {
	fmt.Printf("Drawing a rectangle with width: %v and height: %v\n", r.Width, r.Height)
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t *Triangle) Clone() Shape {
	return &Triangle{
		Base:   t.Base,
		Height: t.Height,
	}
}

func (t *Triangle) Draw() {
	fmt.Printf("Drawing a triangle with base: %v and height: %v\n", t.Base, t.Height)
}

type ShapeDrawer struct {
	Prototypes map[string]Shape
}

func NewShapeDrawer() *ShapeDrawer {
	return &ShapeDrawer{
		Prototypes: make(map[string]Shape),
	}
}

func (s *ShapeDrawer) RegisterPrototype(name string, shape Shape) {
	s.Prototypes[name] = shape
}

func (s *ShapeDrawer) UnregisterPrototype(name string) {
	delete(s.Prototypes, name)
}

func (s *ShapeDrawer) Draw(name string) {
	prototypes, ok := s.Prototypes[name]
	if ok {
		cloneShape := prototypes.Clone()
		cloneShape.Draw()
	}
}
