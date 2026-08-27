package nook

// Resident represents an Animal Crossing character that performs special roles within the player's town, city, or island.
// Characters that are Residents cannot be invited to live within the player's world, but inhabit the world indirectly.
// Residents often serve as merchants, administrators, or quest givers.
type Resident struct {
	Character

	// Roles contains the resident's canonical special-character roles in
	// deterministic key order.
	Roles []Role
}

// HasRole reports whether the resident has the provided special-character
// role.
func (r Resident) HasRole(roleKey Key) bool {
	if roleKey == "" {
		return false
	}

	for _, role := range r.Roles {
		if role.Key == roleKey {
			return true
		}
	}

	return false
}
