package otter

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Lottie.Roles = []nook.Role{role.Proprietor}
		Lyle.Roles = []nook.Role{role.Proprietor}
		Pascal.Roles = []nook.Role{role.RegularVisitor}

		return nook.Residents{
			character.Lottie: Lottie,
			character.Lyle:   Lyle,
			character.Pascal: Pascal,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
