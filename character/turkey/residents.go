package turkey

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Franklin.Roles = []nook.Role{role.HolidayVisitor}

		return nook.Residents{
			character.Franklin: Franklin,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
