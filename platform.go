package nook

// Platform represents a game platform in the Animal Crossing series.
type Platform struct {
	// Key is the language-agnostic key of the platform.
	Key Key

	// Name contains the localized names of the platform.
	Name Languages
}
