package characterabstractfactory

type CharacterType int

const (
	Warrior CharacterType = iota + 1
	Mage
)

func (c CharacterType) String() string {
	return [...]string{"Warior", "Mage"}[c-1]
}

type Character struct {
	CharacterType CharacterType
	Action        string
}

func (c *Character) GetCharacterType() string {
	return c.CharacterType.String()
}

func (c *Character) PerfomAction() string {
	return c.Action
}
