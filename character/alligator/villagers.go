package alligator

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
)

// Villagers is map of villager instances.
var (
	Villagers = nook.Villagers{
		character.Alli:     Alli,
		character.Alfonso:  Alfonso,
		character.Boots:    Boots,
		character.Del:      Del,
		character.Drago:    Drago,
		character.Gayle:    Gayle,
		character.Liz:      Liz,
		character.Pironkon: Pironkon,
		character.Sly:      Sly}
)
