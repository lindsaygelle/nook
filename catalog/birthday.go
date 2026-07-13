package catalog

import (
	"slices"
	"time"

	"github.com/lindsaygelle/nook"
)

// ResidentsByBirthday returns all residents whose birthday exactly matches the
// provided month and day. Results are sorted by animal key and then character
// key for deterministic backend responses.
func ResidentsByBirthday(month time.Month, day uint8) []nook.Resident {
	if day == 0 || month == 0 {
		return nil
	}

	residents := make([]nook.Resident, 0)
	for _, bucket := range AllResidents {
		for _, resident := range bucket {
			if !resident.Birthday.Ok() {
				continue
			}
			if resident.Birthday.Day != day || resident.Birthday.Month != month {
				continue
			}
			residents = append(residents, resident)
		}
	}

	slices.SortFunc(residents, compareResidents)
	return residents
}

// ResidentsByBirthMonth returns all residents whose birthday month exactly
// matches the provided month. Results are sorted by animal key and then
// character key for deterministic backend responses.
func ResidentsByBirthMonth(month time.Month) []nook.Resident {
	if month == 0 {
		return nil
	}

	residents := make([]nook.Resident, 0)
	for _, bucket := range AllResidents {
		for _, resident := range bucket {
			if !resident.Birthday.Ok() {
				continue
			}
			if resident.Birthday.Month != month {
				continue
			}
			residents = append(residents, resident)
		}
	}

	slices.SortFunc(residents, compareResidents)
	return residents
}

// VillagersByBirthday returns all villagers whose birthday exactly matches the
// provided month and day. Results are sorted by animal key and then character
// key for deterministic backend responses.
func VillagersByBirthday(month time.Month, day uint8) []nook.Villager {
	if day == 0 || month == 0 {
		return nil
	}

	villagers := make([]nook.Villager, 0)
	for _, bucket := range AllVillagers {
		for _, villager := range bucket {
			if !villager.Birthday.Ok() {
				continue
			}
			if villager.Birthday.Day != day || villager.Birthday.Month != month {
				continue
			}
			villagers = append(villagers, villager)
		}
	}

	slices.SortFunc(villagers, compareVillagers)
	return villagers
}

// VillagersByBirthMonth returns all villagers whose birthday month exactly
// matches the provided month. Results are sorted by animal key and then
// character key for deterministic backend responses.
func VillagersByBirthMonth(month time.Month) []nook.Villager {
	if month == 0 {
		return nil
	}

	villagers := make([]nook.Villager, 0)
	for _, bucket := range AllVillagers {
		for _, villager := range bucket {
			if !villager.Birthday.Ok() {
				continue
			}
			if villager.Birthday.Month != month {
				continue
			}
			villagers = append(villagers, villager)
		}
	}

	slices.SortFunc(villagers, compareVillagers)
	return villagers
}
