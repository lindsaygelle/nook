package region

import "github.com/lindsaygelle/nook"

var (
	// regions contains the canonical release regions in deterministic key order.
	regions = []nook.Region{
		Australia,
		China,
		Europe,
		Japan,
		NorthAmerica,
		Worldwide,
	}
)

var (
	// regionsByKey contains canonical release regions indexed by key.
	regionsByKey = func() map[nook.Key]nook.Region {
		index := make(map[nook.Key]nook.Region, len(regions))
		for _, region := range regions {
			index[region.Key] = region
		}
		return index
	}()
)

// ByKey returns the canonical release region with the provided key.
func ByKey(key nook.Key) (nook.Region, bool) {
	region, ok := regionsByKey[key]
	return region, ok
}

// List returns all canonical release regions in deterministic key order.
func List() []nook.Region {
	return append([]nook.Region(nil), regions...)
}
