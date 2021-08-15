package nook

import (
	"fmt"

	"golang.org/x/text/language"
)

// Residents is a collection of Resident.
type Residents map[string]Resident

func (v Residents) Add(key language.Tag, resident Resident) {
	name := resident.Name.Must(key)
	if !name.Ok() {
		panic(fmt.Sprintf("%s.Name[%s].Ok() != true", resident.Name.Must(language.AmericanEnglish).Value, key.String()))
	}
	v[name.Value] = resident
}

func (v Residents) Del(key string) bool {
	delete(v, key)
	return v.Has(key)
}

func (v Residents) Each(fn func(string, Resident)) {
	for k, v := range v {
		fn(k, v)
	}
}

func (v Residents) EachWithBreak(fn func(string, Resident) bool) {
	for k, v := range v {
		if fn(k, v) {
			break
		}
	}
}

func (v Residents) Get(key string) (Resident, bool) {
	Resident, ok := v[key]
	return Resident, ok
}

func (v Residents) Has(key string) bool {
	_, ok := v.Get(key)
	return ok
}

func (v Residents) Must(key string) Resident {
	Resident, ok := v.Get(key)
	if !ok {
		panic(fmt.Sprintf("Residents[%s] not found", key))
	}
	return Resident
}
