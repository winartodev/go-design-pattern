package characterabstractfactory

type GoodCharacterFactory struct {
	Character
}

func (c *GoodCharacterFactory) MakeCharacter() ICharacter {
	return &Character{
		CharacterType: Warrior,
		Action:        "Performing Attack and Defense abilities as a Warrior.",
	}
}
