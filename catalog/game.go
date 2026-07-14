package catalog

import (
	"slices"

	"github.com/lindsaygelle/nook"
)

func characterAppearsInGame(character nook.Character, gameKey nook.Key) bool {
	if gameKey == "" {
		return false
	}

	for _, appearance := range character.Games {
		if appearance.Key == gameKey {
			return true
		}
	}
	return false
}

func compareGamesByReleaseOrder(a, b nook.Game) int {
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

func sortedCharacterGames(games []nook.Game) []nook.Game {
	out := append([]nook.Game(nil), games...)
	slices.SortFunc(out, compareGamesByReleaseOrder)
	return out
}

// CharacterGames returns a character's game appearance history in release
// order.
func CharacterGames(character nook.Character) ([]nook.Game, bool) {
	if character.ID() == "" || character.Games == nil {
		return nil, false
	}
	return sortedCharacterGames(character.Games), true
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

// ResidentsByGame returns all residents that appear in the provided game.
// Results are sorted by animal key and then character key for deterministic
// backend responses.
func ResidentsByGame(gameKey nook.Key) []nook.Resident {
	if gameKey == "" {
		return nil
	}

	residents := make([]nook.Resident, 0)
	for _, bucket := range AllResidents {
		for _, resident := range bucket {
			if !characterAppearsInGame(resident.Character, gameKey) {
				continue
			}
			residents = append(residents, resident)
		}
	}

	slices.SortFunc(residents, compareResidents)
	return residents
}

// VillagersByGame returns all villagers that appear in the provided game.
// Results are sorted by animal key and then character key for deterministic
// backend responses.
func VillagersByGame(gameKey nook.Key) []nook.Villager {
	if gameKey == "" {
		return nil
	}

	villagers := make([]nook.Villager, 0)
	for _, bucket := range AllVillagers {
		for _, villager := range bucket {
			if !characterAppearsInGame(villager.Character, gameKey) {
				continue
			}
			villagers = append(villagers, villager)
		}
	}

	slices.SortFunc(villagers, compareVillagers)
	return villagers
}
