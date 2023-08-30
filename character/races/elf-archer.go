package races

import (
	weapon "PATTERNQUEST/items/weapons"
)

type ElfArcher struct {
	weapon weapon.Weapon
}

func (ea *ElfArcher) SetWeapon(w weapon.Weapon) {
	ea.weapon = w
}

func (ea *ElfArcher) GetWeapon() weapon.Weapon {
	return ea.weapon
}

func (ea *ElfArcher) Attack() string {
	return ea.weapon.Attack()
}

func (ea *ElfArcher) Damage() int {
	return ea.weapon.Damage()
}
