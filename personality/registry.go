package personality

import "github.com/lindsaygelle/nook"

var (
	personalities = []nook.Personality{
		BigSister,
		Cranky,
		Jock,
		Lazy,
		Normal,
		Peppy,
		Smug,
		Snooty,
	}
	personalitiesByKey = func() map[nook.Key]nook.Personality {
		index := make(map[nook.Key]nook.Personality, len(personalities))
		for _, personality := range personalities {
			index[personality.Key] = personality
		}
		return index
	}()
)

// ByKey returns the canonical personality with the provided key.
func ByKey(key nook.Key) (nook.Personality, bool) {
	personality, ok := personalitiesByKey[key]
	return personality, ok
}

// List returns all canonical personalities in deterministic key order.
func List() []nook.Personality {
	return append([]nook.Personality(nil), personalities...)
}
