package peacock

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Pave.Roles = []nook.Role{role.HolidayVisitor}

		return nook.Residents{
			character.Pave: Pave,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
