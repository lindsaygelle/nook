package catalog

import "github.com/lindsaygelle/nook"

// ResidentByID returns a resident using an exact global character identifier
// match after normalization.
func ResidentByID(id string) (nook.Resident, bool) {
	query := NormalizeLookupValue(id)
	if query == "" {
		return nook.Resident{}, false
	}

	for _, bucket := range AllResidents {
		for _, resident := range bucket {
			if NormalizeLookupValue(resident.Character.ID().String()) != query {
				continue
			}
			return resident, true
		}
	}
	return nook.Resident{}, false
}

// VillagerByID returns a villager using an exact global character identifier
// match after normalization.
func VillagerByID(id string) (nook.Villager, bool) {
	query := NormalizeLookupValue(id)
	if query == "" {
		return nook.Villager{}, false
	}

	for _, bucket := range AllVillagers {
		for _, villager := range bucket {
			if NormalizeLookupValue(villager.Character.ID().String()) != query {
				continue
			}
			return villager, true
		}
	}
	return nook.Villager{}, false
}
