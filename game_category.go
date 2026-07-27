package nook

// GameCategory represents a series classification for an Animal Crossing game.
type GameCategory struct {
	// Key is the language-agnostic key of the game category.
	Key Key

	// Name contains the localized names of the game category.
	Name Languages
}
