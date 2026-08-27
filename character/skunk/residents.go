package skunk

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Kicks.Roles = []nook.Role{role.Proprietor}

		return nook.Residents{
			character.Kicks: Kicks,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
