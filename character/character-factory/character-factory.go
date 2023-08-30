package character_factory

import "PATTERNQUEST/character"

// Abstract Factory pattern
type CharacterFactory interface {
	CreateArcher() character.Character
	CreateMage() character.Character
	CreateWarrior() character.Character
}
