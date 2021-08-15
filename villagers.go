package nook

import (
	"fmt"

	"golang.org/x/text/language"
)

// Villagers is a collection of Villager.
type Villagers map[string]Villager

func (v Villagers) Add(key language.Tag, villager Villager) {
	name := villager.Name.Must(key)
	if !name.Ok() {
		panic(fmt.Sprintf("%s.Name[%s].Ok() != true", villager.Name.Must(language.AmericanEnglish).Value, key.String()))
	}
	v[name.Value] = villager
}

func (v Villagers) Del(key string) bool {
	delete(v, key)
	return v.Has(key)
}

func (v Villagers) Each(fn func(string, Villager)) {
	for k, v := range v {
		fn(k, v)
	}
}

func (v Villagers) EachWithBreak(fn func(string, Villager) bool) {
	for k, v := range v {
		if fn(k, v) {
			break
		}
	}
}

func (v Villagers) Get(key string) (Villager, bool) {
	villager, ok := v[key]
	return villager, ok
}

func (v Villagers) Has(key string) bool {
	_, ok := v.Get(key)
	return ok
}

func (v Villagers) Must(key string) Villager {
	villager, ok := v.Get(key)
	if !ok {
		panic(fmt.Sprintf("Villagers[%s] not found", key))
	}
	return villager
}
