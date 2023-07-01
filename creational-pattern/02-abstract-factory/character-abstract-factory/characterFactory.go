package characterabstractfactory

type ICharacterFactory interface {
	MakeCharacter() ICharacter
}

func GetCharacterFactory(characterType CharacterType) (ICharacterFactory, error) {
	if characterType == Warrior {
		return &GoodCharacterFactory{}, nil
	} else if characterType == Mage {
		return &EvilCharacterFactory{}, nil
	}

	return nil, nil
}
