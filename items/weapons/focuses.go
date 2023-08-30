package weapon

import spell "PATTERNQUEST/spells"

type Focus interface {
	Weapon
	SetSpell(spell.Spell)
}

type Fetish struct {
	s spell.Spell
}

func (f *Fetish) Attack() string {
	return "Casting Cruel Fist!" // This should be dynamic
}

func (f *Fetish) Damage() int {
	return f.s.Damage()
}

func (f *Fetish) SetSpell(s spell.Spell) {
	f.s = s
}

type Talisman struct {
	s spell.Spell
}

func (t *Talisman) Attack() string {
	return "Casting Fairie Fire!" // This should be dynamic
}

func (t *Talisman) Damage() int {
	return t.s.Damage()
}

func (t *Talisman) SetSpell(s spell.Spell) {
	t.s = s
}
