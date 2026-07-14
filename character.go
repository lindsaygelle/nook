package nook

import "slices"

const characterIDSeparator = ":"

func compareGamesByReleaseOrder(a, b Game) int {
	switch {
	case a.ReleaseOrder == 0 && b.ReleaseOrder != 0:
		return 1
	case a.ReleaseOrder != 0 && b.ReleaseOrder == 0:
		return -1
	case a.ReleaseOrder < b.ReleaseOrder:
		return -1
	case a.ReleaseOrder > b.ReleaseOrder:
		return 1
	}

	switch {
	case a.Key < b.Key:
		return -1
	case a.Key > b.Key:
		return 1
	default:
		return 0
	}
}

// Character is a composite type that combines various attributes of an Animal Crossing character.
type Character struct {
	// Animal represents the animal type of the character.
	Animal

	// Birthday contains the birthday information of the character.
	Birthday Birthday

	// Code is a unique identifier for the character.
	Code Code

	// Games contains the character's game appearances.
	Games []Game

	// Gender represents the gender of the character.
	Gender Gender

	// Key is the character's canonical package key.
	// Keys are not guaranteed to be globally unique across animal types.
	Key Key

	// Name contains the names of the character in different languages.
	Name Languages

	// Special indicates whether the character is special or has a unique role.
	Special bool
}

// AppearsInGame reports whether the character appears in the provided game.
func (c Character) AppearsInGame(gameKey Key) bool {
	if gameKey == "" {
		return false
	}

	for _, game := range c.Games {
		if game.Key == gameKey {
			return true
		}
	}

	return false
}

// FirstGame returns the earliest known game appearance for the character.
func (c Character) FirstGame() (Game, bool) {
	games := c.GamesByReleaseOrder()
	if len(games) == 0 {
		return Game{}, false
	}
	return games[0], true
}

// GameByKey returns the character's game appearance for the provided game key.
func (c Character) GameByKey(gameKey Key) (Game, bool) {
	if gameKey == "" {
		return Game{}, false
	}

	for _, game := range c.Games {
		if game.Key == gameKey {
			return game, true
		}
	}

	return Game{}, false
}

// GamesByReleaseOrder returns the character's game appearances sorted into
// deterministic release order.
func (c Character) GamesByReleaseOrder() []Game {
	games := append([]Game(nil), c.Games...)
	slices.SortFunc(games, compareGamesByReleaseOrder)
	return games
}

// ID returns a globally unique identifier composed from the character's animal
// key and character key.
func (c Character) ID() Key {
	if c.Animal.Key == "" || c.Key == "" {
		return ""
	}
	return Key(string(c.Animal.Key) + characterIDSeparator + string(c.Key))
}

// LastGame returns the latest known game appearance for the character.
func (c Character) LastGame() (Game, bool) {
	games := c.GamesByReleaseOrder()
	if len(games) == 0 {
		return Game{}, false
	}
	return games[len(games)-1], true
}
