package nook

import "golang.org/x/text/language"

type Villagers map[string]Villager

func (v Villagers) Add(language language.Tag, villager Villager) {
	name := villager.Name.Must(language)
	if !name.Ok() {
		panic(villager)
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
		panic(key)
	}
	return villager
}
