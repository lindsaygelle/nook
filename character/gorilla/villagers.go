package gorilla

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
)

// Villagers is map of villager instances.
var (
	// Villagers represents villagers.
	Villagers = nook.Villagers{
		character.Al:     Al,
		character.Boone:  Boone,
		character.Boyd:   Boyd,
		character.Cesar:  Cesar,
		character.Hans:   Hans,
		character.Jane:   Jane,
		character.Louie:  Louie,
		character.Peewee: Peewee,
		character.Rilla:  Rilla,
		character.Rocket: Rocket,
		character.Violet: Violet,
		character.Yodel:  Yodel}
)
