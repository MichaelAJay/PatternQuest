package spell

type FairieFire struct{}

func (ff *FairieFire) Cost() int {
	return 1
}

func (ff *FairieFire) Damage() int {
	return 2
}
