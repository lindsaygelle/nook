package alpaca

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
)

var (
	// Residents is map of Resident instances.
	Residents = nook.Residents{
		character.Cyrus: Cyrus,
		character.Reese: Reese}
)
