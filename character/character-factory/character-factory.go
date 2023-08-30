package character_factory

import "PATTERNQUEST/character"

type CharacterFactory interface {
	CreateArcher() character.Character
	CreateMage() character.Character
	CreateWarrior() character.Character
}
