package catalog

import "github.com/lindsaygelle/nook"

// ResidentsByAnimal returns the resident catalog bucket for an animal key.
func ResidentsByAnimal(animalKey nook.Key) (nook.Residents, bool) {
	residents, ok := AllResidents[animalKey]
	return residents, ok
}

// VillagersByAnimal returns the villager catalog bucket for an animal key.
func VillagersByAnimal(animalKey nook.Key) (nook.Villagers, bool) {
	villagers, ok := AllVillagers[animalKey]
	return villagers, ok
}

// ResidentByKey returns a resident from the given animal bucket using an exact key match.
func ResidentByKey(animalKey, key nook.Key) (nook.Resident, bool) {
	residents, ok := ResidentsByAnimal(animalKey)
	if !ok {
		return nook.Resident{}, false
	}
	return residents.Get(key)
}

// VillagerByKey returns a villager from the given animal bucket using an exact key match.
func VillagerByKey(animalKey, key nook.Key) (nook.Villager, bool) {
	villagers, ok := VillagersByAnimal(animalKey)
	if !ok {
		return nook.Villager{}, false
	}
	return villagers.Get(key)
}
