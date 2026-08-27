package nook

// Role represents a canonical special-character role in the Animal Crossing
// series.
type Role struct {
	// Key is the language-agnostic key of the role.
	Key Key

	// Name contains the localized names of the role.
	Name Languages
}
