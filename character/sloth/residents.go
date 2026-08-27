package sloth

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Leif.Roles = []nook.Role{role.Proprietor}

		return nook.Residents{
			character.Leif: Leif,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
