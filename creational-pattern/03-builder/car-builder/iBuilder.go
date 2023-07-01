package carbuilder

type IBuilder interface {
	SetModel()
	SetColor()
	SetEngine()
	SetNumberOfDoor()
	Build() Car
}
