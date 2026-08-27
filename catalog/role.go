package catalog

import (
	"slices"

	"github.com/lindsaygelle/nook"
)

// ResidentsByRole returns all residents that have the provided special-
// character role. Results are sorted by animal key and then character key for
// deterministic backend responses.
func ResidentsByRole(roleKey nook.Key) []nook.Resident {
	if roleKey == "" {
		return nil
	}

	residents := make([]nook.Resident, 0)
	for _, bucket := range AllResidents {
		for _, resident := range bucket {
			if !resident.HasRole(roleKey) {
				continue
			}
			residents = append(residents, resident)
		}
	}

	slices.SortFunc(residents, compareResidents)
	return residents
}
