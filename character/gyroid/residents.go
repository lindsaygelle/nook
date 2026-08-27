package gyroid

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Lloid.Roles = []nook.Role{role.Government, role.Islander, role.Proprietor}

		return nook.Residents{
			character.Lloid: Lloid,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
