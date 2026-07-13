package nook

// Game represents a game in the Animal Crossing series.
type Game struct {
	// Key is the language-agnostic key of the game.
	Key Key

	// Name contains the localized names of the game.
	Name Languages
}
