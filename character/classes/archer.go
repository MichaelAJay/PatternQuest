package classes

type Archer struct {
	Ammo string
}

func (a *Archer) Shoot() string {
	return "shoots a" + a.Ammo + "!"
}
