package nook

import (
	"fmt"
)

// Villagers is a collection of Villager.
type Villagers map[Key]Villager

// Add adds a Villager to the collection.
func (v Villagers) Add(villager Villager) {
	v[villager.Key] = villager
}

// Del removes a Villager from Villagers using the argument key.
func (v Villagers) Del(key Key) bool {
	delete(v, key)
	return v.Has(key)
}

// Each performs a for-each loop across Villagers, executing the argument function for the Villager at the current key.
func (v Villagers) Each(fn func(Key, Villager)) {
	for k, v := range v {
		fn(k, v)
	}
}

// Each performs a for-each loop across Villagers, executing the argument function for the Villager at the current key.
// Unlike Villagers.Each it is possible to break out of the loop by returning true.
func (v Villagers) EachWithBreak(fn func(Key, Villager) bool) {
	for k, v := range v {
		if fn(k, v) {
			break
		}
	}
}

// Get returns a Villager from Villagers using the argument key.
// Returns an additional boolean indicating whether the Villager was found.
func (v Villagers) Get(key Key) (Villager, bool) {
	villager, ok := v[key]
	return villager, ok
}

// Has returns a boolean indicating whether the Villager exists for the argument key.
func (v Villagers) Has(key Key) bool {
	_, ok := v.Get(key)
	return ok
}

// Must returns a Villager from Villagers using the argument key.
// Unlike Villagers.Get, Villagers.Must panics on the condition a Villager cannot be retrieved for the given key.
func (v Villagers) Must(key Key) Villager {
	villager, ok := v.Get(key)
	if !ok {
		panic(fmt.Sprintf("Villagers[%s] not found", key))
	}
	return villager
}
