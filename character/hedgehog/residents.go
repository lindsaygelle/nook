package hedgehog

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Label.Roles = []nook.Role{role.Proprietor}
		Mabel.Roles = []nook.Role{role.Proprietor}
		Sable.Roles = []nook.Role{role.Proprietor}

		return nook.Residents{
			character.Label: Label,
			character.Mabel: Mabel,
			character.Sable: Sable,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
