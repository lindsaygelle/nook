package ostrich

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
)

// Villagers is map of villager instances.
var (
	// Villagers represents villagers.
	Villagers = nook.Villagers{
		character.Blanche:  Blanche,
		character.Cranston: Cranston,
		character.Flora:    Flora,
		character.Gladys:   Gladys,
		character.Nindori:  Nindori,
		character.Julia:    Julia,
		character.Phil:     Phil,
		character.Phoebe:   Phoebe,
		character.Queenie:  Queenie,
		character.Rio:      Rio,
		character.Sandy:    Sandy,
		character.Sprocket: Sprocket}
)
