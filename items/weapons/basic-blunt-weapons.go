package weapon

type Club struct{}

func (c *Club) Attack() string {
	return "Uses club!"
}

func (c *Club) Damage() int {
	return 2
}
