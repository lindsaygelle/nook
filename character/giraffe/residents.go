package giraffe

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Gracie.Roles = []nook.Role{role.Proprietor}

		return nook.Residents{
			character.Gracie: Gracie,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
