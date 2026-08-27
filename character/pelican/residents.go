package pelican

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Pelly.Roles = []nook.Role{role.Government}
		Pete.Roles = []nook.Role{role.Government}
		Phyllis.Roles = []nook.Role{role.Government}

		return nook.Residents{
			character.Pelly:   Pelly,
			character.Pete:    Pete,
			character.Phyllis: Phyllis,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
