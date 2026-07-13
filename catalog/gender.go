package catalog

import (
	"slices"

	"github.com/lindsaygelle/nook"
)

// ResidentsByGender returns all residents whose gender exactly matches the
// provided gender key. Results are sorted by animal key and then character key
// for deterministic backend responses.
func ResidentsByGender(genderKey nook.Key) []nook.Resident {
	if genderKey == "" {
		return nil
	}

	residents := make([]nook.Resident, 0)
	for _, bucket := range AllResidents {
		for _, resident := range bucket {
			if resident.Character.Gender.Key != genderKey {
				continue
			}
			residents = append(residents, resident)
		}
	}

	slices.SortFunc(residents, compareResidents)
	return residents
}

// VillagersByGender returns all villagers whose gender exactly matches the
// provided gender key. Results are sorted by animal key and then character key
// for deterministic backend responses.
func VillagersByGender(genderKey nook.Key) []nook.Villager {
	if genderKey == "" {
		return nil
	}

	villagers := make([]nook.Villager, 0)
	for _, bucket := range AllVillagers {
		for _, villager := range bucket {
			if villager.Character.Gender.Key != genderKey {
				continue
			}
			villagers = append(villagers, villager)
		}
	}

	slices.SortFunc(villagers, compareVillagers)
	return villagers
}
