package nook

import (
	"fmt"

	"golang.org/x/text/language"
)

// Villagers is a collection of Villager.
type Villagers map[string]Villager

// Add adds a Villager to the collection. Add uses the language.Tag as the key to be retrieved from the Villager.
// On the condition the language.Tag cannot be found in the Villager.Name or the value is not OK, Villagers.Add panics.
// When Villagers.Add panics, the error raised contains the name of the Villager in American English, as this is a known value.
func (v Villagers) Add(key language.Tag, villager Villager) {
	name := villager.Name.Must(key)
	if !name.Ok() {
		panic(fmt.Sprintf("%s.Name[%s].Ok() != true", villager.Name.Must(language.AmericanEnglish).Value, key.String()))
	}
	v[name.Value] = villager
}

// Del removes a Villager from Villagers using the argument key.
func (v Villagers) Del(key string) bool {
	delete(v, key)
	return v.Has(key)
}

// Each performs a for-each loop across Villagers, executing the argument function for the Villager at the current key.
func (v Villagers) Each(fn func(string, Villager)) {
	for k, v := range v {
		fn(k, v)
	}
}

// Each performs a for-each loop across Villagers, executing the argument function for the Villager at the current key.
// Unlike Villagers.Each it is possible to break out of the loop by returning true as a callback.
func (v Villagers) EachWithBreak(fn func(string, Villager) bool) {
	for k, v := range v {
		if fn(k, v) {
			break
		}
	}
}

// Get returns a Villager from Villagers using the argument key.
// Returns an additional boolean indicating whether the Villager was found.
func (v Villagers) Get(key string) (Villager, bool) {
	villager, ok := v[key]
	return villager, ok
}

// Has returns a boolean indicating whether the Villager exists for the argument key.
func (v Villagers) Has(key string) bool {
	_, ok := v.Get(key)
	return ok
}

// Must returns a Villager from Villagers using the argument key.
// Unlike Villagers.Get, Villagers.Must panics on the condition a Villager cannot be retrieved for the given key.
func (v Villagers) Must(key string) Villager {
	villager, ok := v.Get(key)
	if !ok {
		panic(fmt.Sprintf("Villagers[%s] not found", key))
	}
	return villager
}
