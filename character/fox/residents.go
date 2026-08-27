package fox

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Redd.Roles = []nook.Role{role.Proprietor}

		return nook.Residents{
			character.Redd: Redd,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
