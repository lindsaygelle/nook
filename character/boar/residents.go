package boar

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		DaisyMae.Roles = []nook.Role{role.RegularVisitor}
		Joan.Roles = []nook.Role{role.RegularVisitor}

		return nook.Residents{
			character.DaisyMae: DaisyMae,
			character.Joan:     Joan,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
