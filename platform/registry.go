package platform

import "github.com/lindsaygelle/nook"

var (
	// platforms contains the canonical game platforms in deterministic key order.
	platforms = []nook.Platform{
		Android,
		GameCube,
		IOS,
		IQuePlayer,
		Nintendo3DS,
		Nintendo64,
		NintendoDS,
		NintendoSwitch,
		Wii,
		WiiU,
	}
)

var (
	// platformsByKey contains canonical game platforms indexed by key.
	platformsByKey = func() map[nook.Key]nook.Platform {
		index := make(map[nook.Key]nook.Platform, len(platforms))
		for _, platform := range platforms {
			index[platform.Key] = platform
		}
		return index
	}()
)

// ByKey returns the canonical game platform with the provided key.
func ByKey(key nook.Key) (nook.Platform, bool) {
	platform, ok := platformsByKey[key]
	return platform, ok
}

// List returns all canonical game platforms in deterministic key order.
func List() []nook.Platform {
	return append([]nook.Platform(nil), platforms...)
}
