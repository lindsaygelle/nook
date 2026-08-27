package squirrel

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Shaki.Roles = []nook.Role{role.Unused}

		return nook.Residents{
			character.Shaki: Shaki,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
