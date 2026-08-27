package panther

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Katrina.Roles = []nook.Role{role.Proprietor}

		return nook.Residents{
			character.Katrina: Katrina,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
