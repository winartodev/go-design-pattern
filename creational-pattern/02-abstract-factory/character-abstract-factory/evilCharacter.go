package characterabstractfactory

type EvilCharacterFactory struct {
}

func (c *EvilCharacterFactory) MakeCharacter() ICharacter {
	return &Character{
		CharacterType: Mage,
		Action:        "Performing Spellcasting and Teleportation abilities as a Mage.",
	}
}
