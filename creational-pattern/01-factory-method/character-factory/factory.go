package characterfactory

import "fmt"

func CharacterFactory(characterType string) (ICharacter, error) {
	if characterType == "mage" {
		return MageFactory(), nil
	}

	if characterType == "warior" {
		return WariorFactory(), nil
	}

	return nil, fmt.Errorf("wrong character type")
}
