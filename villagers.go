package nook

import (
	"fmt"
)

// Villagers is a collection of Villager characters, where each Villager is identified by a unique Key.
type Villagers map[Key]Villager

// Add adds a Villager to the collection.
func (v Villagers) Add(villager Villager) {
	v[villager.Key] = villager
}

// Del removes a Villager from Villagers using the specified Key.
// It returns true if the Villager was found and removed, otherwise false.
func (v Villagers) Del(key Key) bool {
	_, found := v[key]
	delete(v, key)
	return found
}

// Each iterates over each Villager in the collection, executing the specified function for each one.
func (v Villagers) Each(fn func(Key, Villager)) {
	for k, v := range v {
		fn(k, v)
	}
}

// EachWithBreak iterates over each Villager in the collection, executing the specified function for each one.
// It is possible to break out of the loop by returning true as a callback.
func (v Villagers) EachWithBreak(fn func(Key, Villager) bool) {
	for k, v := range v {
		if fn(k, v) {
			break
		}
	}
}

// Get retrieves a Villager from Villagers using the specified Key.
// It returns the Villager and a boolean indicating whether the Villager was found.
func (v Villagers) Get(key Key) (Villager, bool) {
	villager, ok := v[key]
	return villager, ok
}

// Has checks if a Villager exists for the specified Key in the collection.
// It returns true if the Villager is found, otherwise false.
func (v Villagers) Has(key Key) bool {
	_, ok := v.Get(key)
	return ok
}

// Must retrieves a Villager from Villagers using the specified Key.
// Unlike Villagers.Get, Villagers.Must panics if a Villager cannot be retrieved for the given Key.
func (v Villagers) Must(key Key) Villager {
	villager, ok := v.Get(key)
	if !ok {
		panic(fmt.Sprintf("Villagers[%s] not found", key))
	}
	return villager
}
