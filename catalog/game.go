package catalog

import "github.com/lindsaygelle/nook"

// CharacterGames returns a character's game appearance history in release
// order.
func CharacterGames(character nook.Character) ([]nook.Game, bool) {
	if character.ID() == "" || character.Games == nil {
		return nil, false
	}
	return append([]nook.Game(nil), character.Games...), true
}

// CharacterGamesByID returns a character's game appearance history using an
// exact global character identifier match after normalization.
func CharacterGamesByID(id string) ([]nook.Game, bool) {
	if resident, ok := ResidentByID(id); ok {
		return CharacterGames(resident.Character)
	}
	if villager, ok := VillagerByID(id); ok {
		return CharacterGames(villager.Character)
	}
	return nil, false
}

// FirstAppearance returns the earliest known game appearance for a character.
func FirstAppearance(character nook.Character) (nook.Game, bool) {
	return FirstAppearanceByID(string(character.ID()))
}

// FirstAppearanceByID returns the earliest known game appearance using an
// exact global character identifier match after normalization.
func FirstAppearanceByID(id string) (nook.Game, bool) {
	games, ok := CharacterGamesByID(id)
	if !ok || len(games) == 0 {
		return nook.Game{}, false
	}
	return games[0], true
}

// LastAppearance returns the latest known game appearance for a character.
func LastAppearance(character nook.Character) (nook.Game, bool) {
	return LastAppearanceByID(string(character.ID()))
}

// LastAppearanceByID returns the latest known game appearance using an exact
// global character identifier match after normalization.
func LastAppearanceByID(id string) (nook.Game, bool) {
	games, ok := CharacterGamesByID(id)
	if !ok || len(games) == 0 {
		return nook.Game{}, false
	}
	return games[len(games)-1], true
}
