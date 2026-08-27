package rabbit

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		ZipperTBunny.Roles = []nook.Role{role.HolidayVisitor}

		return nook.Residents{
			character.ZipperTBunny: ZipperTBunny,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
