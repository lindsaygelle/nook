package chameleon

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Flick.Roles = []nook.Role{role.RegularVisitor}
		Nat.Roles = []nook.Role{role.RegularVisitor}

		return nook.Residents{
			character.Flick: Flick,
			character.Nat:   Nat,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
