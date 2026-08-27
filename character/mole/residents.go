package mole

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		DonResetti.Roles = []nook.Role{role.RegularVisitor}
		MrResetti.Roles = []nook.Role{role.RegularVisitor}

		return nook.Residents{
			character.DonResetti: DonResetti,
			character.MrResetti:  MrResetti,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
