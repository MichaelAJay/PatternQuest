package races

import (
	weapon "PATTERNQUEST/items/weapons"
)

type ElfWarrior struct {
	weapon weapon.Weapon
}

func (ew *ElfWarrior) SetWeapon(w weapon.Weapon) {
	ew.weapon = w
}

func (ew *ElfWarrior) GetWeapon() weapon.Weapon {
	return ew.weapon
}

func (ew *ElfWarrior) Attack() string {
	return ew.weapon.Attack()
}

func (ew *ElfWarrior) Damage() int {
	return ew.weapon.Damage()
}
