package camel

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Saharah.Roles = []nook.Role{role.RegularVisitor}

		return nook.Residents{
			character.Saharah: Saharah,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
