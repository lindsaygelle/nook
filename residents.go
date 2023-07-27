package nook

import (
	"fmt"
)

// Residents is a collection of Resident.
type Residents map[Key]Resident

// Add adds a Resident to the collection.
func (v Residents) Add(resident Resident) {
	v[resident.Key] = resident
}

// Del removes a Resident from Residents using the argument key.
func (v Residents) Del(key Key) bool {
	delete(v, key)
	return v.Has(key)
}

// Each performs a for-each loop across Residents, executing the argument function for the Resident at the current key.
func (v Residents) Each(fn func(Key, Resident)) {
	for k, v := range v {
		fn(k, v)
	}
}

// EachWithBreak performs a for-each loop across Residents, executing the argument function for the Resident at the current key.
// Unlike Residents.Each it is possible to break out of the loop by returning true.
func (v Residents) EachWithBreak(fn func(Key, Resident) bool) {
	for k, v := range v {
		if fn(k, v) {
			break
		}
	}
}

// Get returns a Resident from Residents using the argument key.
// Returns an additional boolean indicating whether the Resident was found.
func (v Residents) Get(key Key) (Resident, bool) {
	Resident, ok := v[key]
	return Resident, ok
}

// Has returns a boolean indicating whether the Resident exists for the argument key.
func (v Residents) Has(key Key) bool {
	_, ok := v.Get(key)
	return ok
}

// Must returns a Resident from Residents using the argument key.
// Unlike Residents.Get, Residents.Must panics on the condition a Resident cannot be retrieved for the given key.
func (v Residents) Must(key Key) Resident {
	Resident, ok := v.Get(key)
	if !ok {
		panic(fmt.Sprintf("Residents[%s] not found", key))
	}
	return Resident
}
