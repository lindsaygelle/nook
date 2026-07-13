package nook

// Personality represents an Animal Crossing character's personality type.
type Personality struct {
	// Key is the language-agnostic key of the personality.
	Key Key

	// Name contains the name of the personality type in different languages.
	Name Languages
}
