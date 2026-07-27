package gamecategory

import "github.com/lindsaygelle/nook"

var (
	// categories contains the canonical game categories in deterministic key order.
	categories = []nook.GameCategory{
		Mainline,
		Mobile,
		Spinoff,
	}
)

var (
	// categoriesByKey contains canonical game categories indexed by key.
	categoriesByKey = func() map[nook.Key]nook.GameCategory {
		index := make(map[nook.Key]nook.GameCategory, len(categories))
		for _, category := range categories {
			index[category.Key] = category
		}
		return index
	}()
)

// ByKey returns the canonical game category with the provided key.
func ByKey(key nook.Key) (nook.GameCategory, bool) {
	category, ok := categoriesByKey[key]
	return category, ok
}

// List returns all canonical game categories in deterministic key order.
func List() []nook.GameCategory {
	return append([]nook.GameCategory(nil), categories...)
}
