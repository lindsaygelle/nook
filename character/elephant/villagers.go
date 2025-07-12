package elephant

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
)

// Villagers is map of villager instances.
var (
	// Villagers represents villagers.
	Villagers = nook.Villagers{
		character.Axel:   Axel,
		character.BigTop: BigTop,
		character.Chai:   Chai,
		character.Cyd:    Cyd,
		character.Dizzy:  Dizzy,
		character.Elina:  Elina,
		character.Ellie:  Ellie,
		character.Eloise: Eloise,
		character.Margie: Margie,
		character.Opal:   Opal,
		character.Paolo:  Paolo,
		character.Tucker: Tucker,
		character.Tia:    Tia}
)
