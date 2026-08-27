package walrus

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Wendell.Roles = []nook.Role{role.RegularVisitor}

		return nook.Residents{
			character.Wendell: Wendell,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
