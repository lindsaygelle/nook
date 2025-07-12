package tiger

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
)

// Villagers is map of villager instances.
var (
	// Villagers represents villagers.
	Villagers = nook.Villagers{
		character.Bangle:   Bangle,
		character.Bianca:   Bianca,
		character.Claudia:  Claudia,
		character.Leonardo: Leonardo,
		character.Rolf:     Rolf,
		character.Rowan:    Rowan,
		character.Tybalt:   Tybalt}
)
