package reindeer

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Jingle.Roles = []nook.Role{role.HolidayVisitor}

		return nook.Residents{
			character.Jingle: Jingle,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
