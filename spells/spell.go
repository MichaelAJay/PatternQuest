package spell

type Spell interface {
	Cost() int
	Damage() int
}
