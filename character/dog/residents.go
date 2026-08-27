package dog

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Booker.Roles = []nook.Role{role.Government}
		Copper.Roles = []nook.Role{role.Government}
		Digby.Roles = []nook.Role{role.Proprietor}
		Harriet.Roles = []nook.Role{role.Proprietor}
		Harvey.Roles = []nook.Role{role.Proprietor}
		Isabelle.Roles = []nook.Role{role.Government}
		KKSlider.Roles = []nook.Role{role.RegularVisitor}
		Serena.Roles = []nook.Role{role.RegularVisitor}

		return nook.Residents{
			character.Booker:   Booker,
			character.Copper:   Copper,
			character.Digby:    Digby,
			character.Harriet:  Harriet,
			character.Harvey:   Harvey,
			character.Isabelle: Isabelle,
			character.KKSlider: KKSlider,
			character.Serena:   Serena,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
