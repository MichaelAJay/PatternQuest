package spell

type CruelFist struct{}

func (cf *CruelFist) Cost() int {
	return 2
}

func (cf *CruelFist) Damage() int {
	return 2
}
