package gender

import "github.com/lindsaygelle/nook"

var (
	genders = []nook.Gender{
		Female,
		Male,
	}
	gendersByKey = func() map[nook.Key]nook.Gender {
		index := make(map[nook.Key]nook.Gender, len(genders))
		for _, gender := range genders {
			index[gender.Key] = gender
		}
		return index
	}()
)

// ByKey returns the canonical gender with the provided key.
func ByKey(key nook.Key) (nook.Gender, bool) {
	gender, ok := gendersByKey[key]
	return gender, ok
}

// List returns all canonical genders in deterministic key order.
func List() []nook.Gender {
	return append([]nook.Gender(nil), genders...)
}
