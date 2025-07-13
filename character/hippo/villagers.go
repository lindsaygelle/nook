package hippo

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
)

// Villagers is map of villager instances.
var (
	// Villagers represents villagers.
	Villagers = nook.Villagers{
		character.Bitty:   Bitty,
		character.Biff:    Biff,
		character.Bertha:  Bertha,
		character.Bubbles: Bubbles,
		character.Clara:   Clara,
		character.Hippeux: Hippeux,
		character.Harry:   Harry,
		character.Lulu:    Lulu,
		character.Rocco:   Rocco,
		character.Rollo:   Rollo}
)
