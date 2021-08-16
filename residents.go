package nook

import (
	"fmt"

	"golang.org/x/text/language"
)

// Residents is a collection of Resident.
type Residents map[string]Resident

// Add adds a Resident to the collection. Add uses the language.Tag as the key to be retrieved from the Resident.
// On the condiiton the language.Tag cannot be found in the Resident.Name or the value is not OK, Residents.Add panics.
// When Residents.Add panics, the error raised contains the name of the Resident in American English, as this is a known value.
func (v Residents) Add(key language.Tag, resident Resident) {
	name := resident.Name.Must(key)
	if !name.Ok() {
		panic(fmt.Sprintf("%s.Name[%s].Ok() != true", resident.Name.Must(language.AmericanEnglish).Value, key.String()))
	}
	v[name.Value] = resident
}

// Del removes a Resident from Residents using the argument key.
func (v Residents) Del(key string) bool {
	delete(v, key)
	return v.Has(key)
}

// Each performs a for-each loop across Residents, executing the argument function for the Resident at the current key.
func (v Residents) Each(fn func(string, Resident)) {
	for k, v := range v {
		fn(k, v)
	}
}

// Each performs a for-each loop across Residents, executing the argument function for the Resident at the current key.
// Unlike Residents.Each it is possible to break out of the loop by returning true as a callback.
func (v Residents) EachWithBreak(fn func(string, Resident) bool) {
	for k, v := range v {
		if fn(k, v) {
			break
		}
	}
}

// Get returns a Resident from Residents using the argument key.
// Returns an additional boolean indicating whether the Resident was found.
func (v Residents) Get(key string) (Resident, bool) {
	Resident, ok := v[key]
	return Resident, ok
}

// Has returns a boolean indicating whether the Resident exists for the argument key.
func (v Residents) Has(key string) bool {
	_, ok := v.Get(key)
	return ok
}

// Must returns a Resident from Residents using the argument key.
// Unlike Residents.Get, Residents.Must panics on the condiiton a Resident cannot be retrieved for the given key.
func (v Residents) Must(key string) Resident {
	Resident, ok := v.Get(key)
	if !ok {
		panic(fmt.Sprintf("Residents[%s] not found", key))
	}
	return Resident
}
