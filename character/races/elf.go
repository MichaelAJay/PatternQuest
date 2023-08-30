package races

import (
	"PATTERNQUEST/character"
	weapon "PATTERNQUEST/items/weapons"
)

type ElfFactory struct{}

func (e *ElfFactory) CreateArcher() character.Character {
	archer := &ElfArcher{}
	archer.SetWeapon(&weapon.Arrow{})
	return archer
}

func (e *ElfFactory) CreateMage() character.Character {
	mage := &ElfMage{}
	mage.SetWeapon(&weapon.Talisman{})
	return mage
}

func (e *ElfFactory) CreateWarrior() character.Character {
	warrior := &ElfWarrior{}
	warrior.SetWeapon(&weapon.Longsword{})
	return warrior
}
