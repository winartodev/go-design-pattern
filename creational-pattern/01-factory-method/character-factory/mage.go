package characterfactory

type Mage struct {
	Character
}

func MageFactory() ICharacter {
	return &Mage{
		Character{
			Name: "John",
		},
	}
}

func (m *Mage) PerformAbility() (ability string) {
	return "Performing Spellcasting and Teleportation abilities as a Mage."
}
