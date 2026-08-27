package tortoise

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Cornimer.Roles = []nook.Role{role.HolidayVisitor}
		Tortimer.Roles = []nook.Role{role.Government, role.Islander}

		return nook.Residents{
			character.Cornimer: Cornimer,
			character.Tortimer: Tortimer,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
