package tapir

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Luna.Roles = []nook.Role{role.Proprietor}

		return nook.Residents{
			character.Luna: Luna,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
