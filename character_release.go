package nook

// CharacterRelease represents a single known game release entry in a
// character's appearance history.
type CharacterRelease struct {
	// Game is the game associated with the release entry.
	Game Game

	// ReleaseDate is the regional release date associated with the game entry.
	ReleaseDate ReleaseDate
}
