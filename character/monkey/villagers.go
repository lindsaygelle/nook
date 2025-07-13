package monkey

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
)

// Villagers is map of villager instances.
var (
	// Villagers represents villagers.
	Villagers = nook.Villagers{
		character.Champ: Champ,
		character.Deli:  Deli,
		character.Elise: Elise,
		character.Flip:  Flip,
		character.Monty: Monty,
		character.Nana:  Nana,
		character.Shari: Shari,
		character.Simon: Simon,
		character.Tammi: Tammi}
)
