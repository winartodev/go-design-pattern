package characterfactory

type Warior struct {
	Character
}

func WariorFactory() ICharacter {
	return &Warior{
		Character{
			Name: "Fade",
		},
	}
}

func (m *Warior) PerformAbility() (ability string) {
	return "Performing Attack and Defense abilities as a Warrior."
}
