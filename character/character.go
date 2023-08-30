package character

import weapon "PATTERNQUEST/items/weapons"

type Character interface {
	Attack() string
	Damage() int
	SetWeapon(weapon.Weapon)
}
