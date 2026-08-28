package catalog

import (
	"slices"

	"github.com/lindsaygelle/nook"
)

// ResidentsByZodiacSign returns all residents with the provided zodiac sign.
// Results are sorted by animal key and then character key for deterministic
// backend responses.
func ResidentsByZodiacSign(zodiacSignKey nook.Key) []nook.Resident {
	if zodiacSignKey == "" {
		return nil
	}

	residents := make([]nook.Resident, 0)
	for _, bucket := range AllResidents {
		for _, resident := range bucket {
			if !resident.Character.HasZodiacSign(zodiacSignKey) {
				continue
			}
			residents = append(residents, resident)
		}
	}

	slices.SortFunc(residents, compareResidents)
	return residents
}

// VillagersByZodiacSign returns all villagers with the provided zodiac sign.
// Results are sorted by animal key and then character key for deterministic
// backend responses.
func VillagersByZodiacSign(zodiacSignKey nook.Key) []nook.Villager {
	if zodiacSignKey == "" {
		return nil
	}

	villagers := make([]nook.Villager, 0)
	for _, bucket := range AllVillagers {
		for _, villager := range bucket {
			if !villager.Character.HasZodiacSign(zodiacSignKey) {
				continue
			}
			villagers = append(villagers, villager)
		}
	}

	slices.SortFunc(villagers, compareVillagers)
	return villagers
}
