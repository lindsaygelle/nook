package raccoon

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		TomNook.Roles = []nook.Role{role.Proprietor}

		return nook.Residents{
			character.TomNook: TomNook,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
