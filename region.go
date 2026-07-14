package nook

// Region represents a major release region in the Animal Crossing series.
type Region struct {
	// Key is the language-agnostic key of the region.
	Key Key

	// Name contains the localized names of the region.
	Name Languages
}
