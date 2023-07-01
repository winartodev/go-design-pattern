package characterfactory

type Character struct {
	Name string
}

func (c *Character) GetName() (name string) {
	return c.Name
}

func (c *Character) SetName(name string) {
	c.Name = name
}

func (c *Character) PerformAbility() (ability string) {
	return ""
}
