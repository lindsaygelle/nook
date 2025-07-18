package cow

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
)

// Villagers is map of villager instances.
var (
	// Villagers represents villagers.
	Villagers = nook.Villagers{
		character.Belle:   Belle,
		character.Bessie:  Bessie,
		character.Carrot:  Carrot,
		character.Naomi:   Naomi,
		character.Norma:   Norma,
		character.Patty:   Patty,
		character.Petunia: Petunia,
		character.Tipper:  Tipper}
)
