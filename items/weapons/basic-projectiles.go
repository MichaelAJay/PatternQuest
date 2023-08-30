package weapon

type Stone struct{}
type Arrow struct{}
type Bolt struct{}

func (s *Stone) Attack() string {
	return "Throw rock!"
}

func (s *Stone) Damage() int {
	return 1
}

func (a *Arrow) Attack() string {
	return "Arrow shoot!"
}

func (a *Arrow) Damage() int {
	return 2
}

func (b *Bolt) Attack() string {
	return "Bolt shoot!"
}

func (b *Bolt) Damage() int {
	return 3
}
