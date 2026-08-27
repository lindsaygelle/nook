package bird

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Beppe.Roles = []nook.Role{role.Proprietor}
		Carlo.Roles = []nook.Role{role.Proprietor}
		Giovanni.Roles = []nook.Role{role.Proprietor}

		return nook.Residents{
			character.Beppe:    Beppe,
			character.Carlo:    Carlo,
			character.Giovanni: Giovanni,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
