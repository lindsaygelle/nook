package nook

// Gender represents the gender of an Animal Crossing character.
type Gender struct {
	// Key is the language-agnostic key of the gender.
	Key Key

	// Name contains the gender name of the character in different languages.
	Name Languages
}
