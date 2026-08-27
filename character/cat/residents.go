package cat

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

var (
	// residents contains canonical residents with role metadata.
	residents = func() nook.Residents {
		Blanca.Roles = []nook.Role{role.RegularVisitor}
		Kaitlin.Roles = []nook.Role{role.RegularVisitor}
		Katie.Roles = []nook.Role{role.RegularVisitor}
		Rover.Roles = []nook.Role{role.RegularVisitor}

		return nook.Residents{
			character.Blanca:  Blanca,
			character.Kaitlin: Kaitlin,
			character.Katie:   Katie,
			character.Rover:   Rover,
		}
	}()
)

var (
	// Residents represents residents.
	Residents = residents
)
