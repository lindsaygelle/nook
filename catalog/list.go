package catalog

import (
	"slices"
	"strings"

	"github.com/lindsaygelle/nook"
)

func sortedResidents(residents nook.Residents) []nook.Resident {
	out := make([]nook.Resident, 0, len(residents))
	for _, resident := range residents {
		out = append(out, resident)
	}
	slices.SortFunc(out, compareResidents)
	return out
}

func sortedVillagers(villagers nook.Villagers) []nook.Villager {
	out := make([]nook.Villager, 0, len(villagers))
	for _, villager := range villagers {
		out = append(out, villager)
	}
	slices.SortFunc(out, compareVillagers)
	return out
}

func compareResidents(a, b nook.Resident) int {
	return compareCharacters(a.Character, b.Character)
}

func compareVillagers(a, b nook.Villager) int {
	return compareCharacters(a.Character, b.Character)
}

func compareCharacters(a, b nook.Character) int {
	if a.Animal.Key != b.Animal.Key {
		return strings.Compare(string(a.Animal.Key), string(b.Animal.Key))
	}
	return strings.Compare(string(a.Key), string(b.Key))
}

// ResidentListByAnimal returns a resident bucket as a deterministically sorted
// slice ordered by animal key and then character key.
func ResidentListByAnimal(animalKey nook.Key) ([]nook.Resident, bool) {
	residents, ok := ResidentsByAnimal(animalKey)
	if !ok {
		return nil, false
	}
	return sortedResidents(residents), true
}

// VillagerListByAnimal returns a villager bucket as a deterministically sorted
// slice ordered by animal key and then character key.
func VillagerListByAnimal(animalKey nook.Key) ([]nook.Villager, bool) {
	villagers, ok := VillagersByAnimal(animalKey)
	if !ok {
		return nil, false
	}
	return sortedVillagers(villagers), true
}
