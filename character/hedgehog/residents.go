package hedgehog

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
)

var (
	// Residents represents residents.
	Residents = nook.Residents{
		character.Label: Label,
		character.Mabel: Mabel,
		character.Sable: Sable}
)
