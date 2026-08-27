package turtle

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Grams.Roles = []nook.Role{role.Islander}
		Kappn.Roles = []nook.Role{role.Islander, role.Proprietor}
		Leila.Roles = []nook.Role{role.Islander}
		Leilani.Roles = []nook.Role{role.Islander}

		return nook.Residents{
			character.Grams:   Grams,
			character.Kappn:   Kappn,
			character.Leila:   Leila,
			character.Leilani: Leilani,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
