package dodo

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Orville.Roles = []nook.Role{role.Government}
		Wilbur.Roles = []nook.Role{role.Government}

		return nook.Residents{
			character.Orville: Orville,
			character.Wilbur:  Wilbur,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
