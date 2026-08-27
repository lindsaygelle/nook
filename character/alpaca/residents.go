package alpaca

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Cyrus.Roles = []nook.Role{role.Proprietor}
		Reese.Roles = []nook.Role{role.Proprietor}

		return nook.Residents{
			character.Cyrus: Cyrus,
			character.Reese: Reese,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
