package role

import "github.com/lindsaygelle/nook"

var (
	// roles contains canonical special-character roles in deterministic key order.
	roles = []nook.Role{
		Government,
		HolidayVisitor,
		Islander,
		Proprietor,
		RegularVisitor,
		Unused,
	}
)

var (
	// rolesByKey contains canonical special-character roles indexed by key.
	rolesByKey = func() map[nook.Key]nook.Role {
		index := make(map[nook.Key]nook.Role, len(roles))
		for _, role := range roles {
			index[role.Key] = role
		}
		return index
	}()
)

// ByKey returns the canonical special-character role with the provided key.
func ByKey(key nook.Key) (nook.Role, bool) {
	role, ok := rolesByKey[key]
	return role, ok
}

// List returns all canonical special-character roles in deterministic key order.
func List() []nook.Role {
	return append([]nook.Role(nil), roles...)
}
