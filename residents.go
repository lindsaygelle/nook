package nook

import (
	"fmt"
)

// Residents is a collection of Resident characters, where each Resident is identified by a unique Key.
type Residents map[Key]Resident

// Add adds a Resident to the collection.
func (v Residents) Add(resident Resident) {
	v[resident.Key] = resident
}

// Del removes a Resident from Residents using the specified Key.
// It returns true if the Resident was found and removed, otherwise false.
func (v Residents) Del(key Key) bool {
	_, found := v[key]
	delete(v, key)
	return found
}

// Each iterates over each Resident in the collection, executing the specified function for each one.
func (v Residents) Each(fn func(Key, Resident)) {
	for k, v := range v {
		fn(k, v)
	}
}

// EachWithBreak iterates over each Resident in the collection, executing the specified function for each one.
// It is possible to break out of the loop by returning true as a callback.
func (v Residents) EachWithBreak(fn func(Key, Resident) bool) {
	for k, v := range v {
		if fn(k, v) {
			break
		}
	}
}

// Get retrieves a Resident from Residents using the specified Key.
// It returns the Resident and a boolean indicating whether the Resident was found.
func (v Residents) Get(key Key) (Resident, bool) {
	resident, ok := v[key]
	return resident, ok
}

// Has checks if a Resident exists for the specified Key in the collection.
// It returns true if the Resident is found, otherwise false.
func (v Residents) Has(key Key) bool {
	_, ok := v.Get(key)
	return ok
}

// Must retrieves a Resident from Residents using the specified Key.
// Unlike Residents.Get, Residents.Must panics if a Resident cannot be retrieved for the given Key.
func (v Residents) Must(key Key) Resident {
	resident, ok := v.Get(key)
	if !ok {
		panic(fmt.Sprintf("Residents[%s] not found", key))
	}
	return resident
}
