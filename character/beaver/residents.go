package beaver

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		CJ.Roles = []nook.Role{role.RegularVisitor}
		Chip.Roles = []nook.Role{role.RegularVisitor}

		return nook.Residents{
			character.CJ:   CJ,
			character.Chip: Chip,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
