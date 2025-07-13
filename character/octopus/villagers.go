package octopus

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
)

// Villagers is map of villager instances.
var (
	// Villagers represents villagers.
	Villagers = nook.Villagers{
		character.Inkwell:  Inkwell,
		character.Marina:   Marina,
		character.Octavian: Octavian,
		character.Zucker:   Zucker}
)
