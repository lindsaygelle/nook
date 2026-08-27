package catalog

import (
	"slices"

	"github.com/lindsaygelle/nook"
)

// CharacterGames returns a character's game appearance history in release
// order.
func CharacterGames(character nook.Character) ([]nook.Game, bool) {
	if character.ID() == "" || character.Games == nil {
		return nil, false
	}
	return character.GamesByReleaseOrder(), true
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
	if resident, ok := ResidentByID(id); ok {
		return resident.Character.FirstGame()
	}
	if villager, ok := VillagerByID(id); ok {
		return villager.Character.FirstGame()
	}
	return nook.Game{}, false
}

// LastAppearance returns the latest known game appearance for a character.
func LastAppearance(character nook.Character) (nook.Game, bool) {
	return LastAppearanceByID(string(character.ID()))
}

// LastAppearanceByID returns the latest known game appearance using an exact
// global character identifier match after normalization.
func LastAppearanceByID(id string) (nook.Game, bool) {
	if resident, ok := ResidentByID(id); ok {
		return resident.Character.LastGame()
	}
	if villager, ok := VillagerByID(id); ok {
		return villager.Character.LastGame()
	}
	return nook.Game{}, false
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
			if !resident.Character.AppearsInGame(gameKey) {
				continue
			}
			residents = append(residents, resident)
		}
	}

	slices.SortFunc(residents, compareResidents)
	return residents
}

// ResidentsByGameCategory returns all residents that appear in at least one
// game within the provided category. Results are sorted by animal key and then
// character key for deterministic backend responses.
func ResidentsByGameCategory(categoryKey nook.Key) []nook.Resident {
	if categoryKey == "" {
		return nil
	}

	residents := make([]nook.Resident, 0)
	for _, bucket := range AllResidents {
		for _, resident := range bucket {
			if len(resident.Character.GamesByCategory(categoryKey)) == 0 {
				continue
			}
			residents = append(residents, resident)
		}
	}

	slices.SortFunc(residents, compareResidents)
	return residents
}

// ResidentsByGamePlatform returns all residents that appear in at least one
// game on the provided platform. Results are sorted by animal key and then
// character key for deterministic backend responses.
func ResidentsByGamePlatform(platformKey nook.Key) []nook.Resident {
	if platformKey == "" {
		return nil
	}

	residents := make([]nook.Resident, 0)
	for _, bucket := range AllResidents {
		for _, resident := range bucket {
			if len(resident.Character.GamesByPlatform(platformKey)) == 0 {
				continue
			}
			residents = append(residents, resident)
		}
	}

	slices.SortFunc(residents, compareResidents)
	return residents
}

// ResidentsByGameRegion returns all residents that appear in at least one game
// released in the provided region. Results are sorted by animal key and then
// character key for deterministic backend responses.
func ResidentsByGameRegion(regionKey nook.Key) []nook.Resident {
	if regionKey == "" {
		return nil
	}

	residents := make([]nook.Resident, 0)
	for _, bucket := range AllResidents {
		for _, resident := range bucket {
			if len(resident.Character.GamesByRegion(regionKey)) == 0 {
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
			if !villager.Character.AppearsInGame(gameKey) {
				continue
			}
			villagers = append(villagers, villager)
		}
	}

	slices.SortFunc(villagers, compareVillagers)
	return villagers
}

// VillagersByGameCategory returns all villagers that appear in at least one
// game within the provided category. Results are sorted by animal key and then
// character key for deterministic backend responses.
func VillagersByGameCategory(categoryKey nook.Key) []nook.Villager {
	if categoryKey == "" {
		return nil
	}

	villagers := make([]nook.Villager, 0)
	for _, bucket := range AllVillagers {
		for _, villager := range bucket {
			if len(villager.Character.GamesByCategory(categoryKey)) == 0 {
				continue
			}
			villagers = append(villagers, villager)
		}
	}

	slices.SortFunc(villagers, compareVillagers)
	return villagers
}

// VillagersByGamePlatform returns all villagers that appear in at least one
// game on the provided platform. Results are sorted by animal key and then
// character key for deterministic backend responses.
func VillagersByGamePlatform(platformKey nook.Key) []nook.Villager {
	if platformKey == "" {
		return nil
	}

	villagers := make([]nook.Villager, 0)
	for _, bucket := range AllVillagers {
		for _, villager := range bucket {
			if len(villager.Character.GamesByPlatform(platformKey)) == 0 {
				continue
			}
			villagers = append(villagers, villager)
		}
	}

	slices.SortFunc(villagers, compareVillagers)
	return villagers
}

// VillagersByGameRegion returns all villagers that appear in at least one game
// released in the provided region. Results are sorted by animal key and then
// character key for deterministic backend responses.
func VillagersByGameRegion(regionKey nook.Key) []nook.Villager {
	if regionKey == "" {
		return nil
	}

	villagers := make([]nook.Villager, 0)
	for _, bucket := range AllVillagers {
		for _, villager := range bucket {
			if len(villager.Character.GamesByRegion(regionKey)) == 0 {
				continue
			}
			villagers = append(villagers, villager)
		}
	}

	slices.SortFunc(villagers, compareVillagers)
	return villagers
}
