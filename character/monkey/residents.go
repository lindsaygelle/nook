package monkey

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Porter.Roles = []nook.Role{role.Government}

		return nook.Residents{
			character.Porter: Porter,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
