package races

import (
	weapon "PATTERNQUEST/items/weapons"
	spell "PATTERNQUEST/spells"
)

type ElfMage struct {
	weapon weapon.Focus
}

func (em *ElfMage) SetWeapon(w weapon.Weapon) {
	if focus, ok := w.(weapon.Focus); ok {
		em.weapon = focus
		em.weapon.SetSpell(&spell.FairieFire{})
	}
}

func (em *ElfMage) GetWeapon() weapon.Focus {
	return em.weapon
}

func (em *ElfMage) Attack() string {
	return em.weapon.Attack()
}

func (em *ElfMage) Damage() int {
	return em.weapon.Damage()
}
