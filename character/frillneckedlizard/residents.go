package frillneckedlizard

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Frillard.Roles = []nook.Role{role.Proprietor}

		return nook.Residents{
			character.Frillard: Frillard,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
