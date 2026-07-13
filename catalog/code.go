package catalog

import "github.com/lindsaygelle/nook"

// ResidentByCode returns a resident using an exact code match after
// normalization. Characters without a populated code are skipped.
func ResidentByCode(code string) (nook.Resident, bool) {
	query := NormalizeLookupValue(code)
	if query == "" {
		return nook.Resident{}, false
	}

	for _, bucket := range AllResidents {
		for _, resident := range bucket {
			if !resident.Code.Ok() {
				continue
			}
			if NormalizeLookupValue(resident.Code.Value) != query {
				continue
			}
			return resident, true
		}
	}
	return nook.Resident{}, false
}

// VillagerByCode returns a villager using an exact code match after
// normalization. Characters without a populated code are skipped.
func VillagerByCode(code string) (nook.Villager, bool) {
	query := NormalizeLookupValue(code)
	if query == "" {
		return nook.Villager{}, false
	}

	for _, bucket := range AllVillagers {
		for _, villager := range bucket {
			if !villager.Code.Ok() {
				continue
			}
			if NormalizeLookupValue(villager.Code.Value) != query {
				continue
			}
			return villager, true
		}
	}
	return nook.Villager{}, false
}
