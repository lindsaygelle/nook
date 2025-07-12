package rhinoceros

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
)

// Villagers is map of villager instances.
var (
	// Villagers represents villagers.
	Villagers = nook.Villagers{
		character.Hornsby:  Hornsby,
		character.Merengue: Merengue,
		character.Patricia: Patricia,
		character.Petunia:  Petunia,
		character.Renee:    Renee,
		character.Rhonda:   Rhonda,
		character.Spike:    Spike,
		character.Tiara:    Tiara,
		character.Tank:     Tank}
)
