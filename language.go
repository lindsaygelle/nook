package nook

import (
	"fmt"

	"golang.org/x/text/language"
)

// Languages is a collection of language information, where each language tag maps to a Name.
type Languages map[language.Tag]Name

// Add adds a Name to the collection. The Name.Language is used as the key to store it in Languages.
func (v Languages) Add(name Name) {
	v[name.Language] = name
}

// Del removes a Name from Languages using the specified language tag as the key.
// It returns true if the Name was found and removed, otherwise false.
func (v Languages) Del(key language.Tag) bool {
	_, found := v[key]
	delete(v, key)
	return found
}

// Each iterates over each language entry in Languages, executing the specified function for each Name.
func (v Languages) Each(fn func(language.Tag, Name)) {
	for k, v := range v {
		fn(k, v)
	}
}

// EachWithBreak iterates over each language entry in Languages, executing the specified function for each Name.
// It is possible to break out of the loop by returning true as a callback.
func (v Languages) EachWithBreak(fn func(language.Tag, Name) bool) {
	for k, v := range v {
		if fn(k, v) {
			break
		}
	}
}

// Get retrieves a Name from Languages using the specified language tag as the key.
// It returns the Name and a boolean indicating whether the Name was found.
func (v Languages) Get(key language.Tag) (Name, bool) {
	name, ok := v[key]
	return name, ok
}

// Has checks if a Name exists for the specified language tag in Languages.
// It returns true if the Name is found, otherwise false.
func (v Languages) Has(key language.Tag) bool {
	_, ok := v.Get(key)
	return ok
}

// Must retrieves a Name from Languages using the specified language tag as the key.
// Unlike Languages.Get, Languages.Must panics if a Name cannot be retrieved for the given key.
func (v Languages) Must(key language.Tag) Name {
	name, ok := v.Get(key)
	if !ok {
		panic(fmt.Sprintf("Languages[%s] not found", key.String()))
	}
	return name
}
