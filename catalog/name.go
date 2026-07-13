package catalog

import (
	"slices"
	"strings"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

// NormalizeLookupValue trims surrounding whitespace and folds case for exact
// lookup operations. It does not remove accents or perform fuzzy matching.
func NormalizeLookupValue(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

// ResidentsByName returns all residents whose localized name matches the
// provided value after normalization. Results are sorted by animal key and then
// character key for deterministic backend responses.
func ResidentsByName(tag language.Tag, name string) []nook.Resident {
	query := NormalizeLookupValue(name)
	if query == "" {
		return nil
	}

	residents := make([]nook.Resident, 0)
	for _, bucket := range AllResidents {
		for _, resident := range bucket {
			localizedName, ok := resident.Name.Get(tag)
			if !ok {
				continue
			}
			if NormalizeLookupValue(localizedName.Value) != query {
				continue
			}
			residents = append(residents, resident)
		}
	}

	slices.SortFunc(residents, compareResidents)
	return residents
}

// VillagersByName returns all villagers whose localized name matches the
// provided value after normalization. Results are sorted by animal key and then
// character key for deterministic backend responses.
func VillagersByName(tag language.Tag, name string) []nook.Villager {
	query := NormalizeLookupValue(name)
	if query == "" {
		return nil
	}

	villagers := make([]nook.Villager, 0)
	for _, bucket := range AllVillagers {
		for _, villager := range bucket {
			localizedName, ok := villager.Name.Get(tag)
			if !ok {
				continue
			}
			if NormalizeLookupValue(localizedName.Value) != query {
				continue
			}
			villagers = append(villagers, villager)
		}
	}

	slices.SortFunc(villagers, compareVillagers)
	return villagers
}
