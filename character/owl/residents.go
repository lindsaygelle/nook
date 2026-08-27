package owl

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Blathers.Roles = []nook.Role{role.Government}
		Celeste.Roles = []nook.Role{role.Government}

		return nook.Residents{
			character.Blathers: Blathers,
			character.Celeste:  Celeste,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
