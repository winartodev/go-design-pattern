package characterfactory

type ICharacter interface {
	GetName() (name string)
	SetName(name string)
	PerformAbility() (ability string)
}
