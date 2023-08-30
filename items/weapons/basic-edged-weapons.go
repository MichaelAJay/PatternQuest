package weapon

type Longsword struct{}

func (l *Longsword) Attack() string {
	return "Uses longsword!"
}

func (l *Longsword) Damage() int {
	return 2
}
