package hamster

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
)

// Villagers is map of villager instances.
var (
	Villagers = nook.Villagers{
		character.Apple:    Apple,
		character.Clay:     Clay,
		character.Flurry:   Flurry,
		character.Graham:   Graham,
		character.Hamphrey: Hamphrey,
		character.Hamlet:   Hamlet,
		character.Holden:   Holden,
		character.Rodney:   Rodney,
		character.Soleil:   Soleil}
)
