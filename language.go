package nook

import (
	"fmt"

	"golang.org/x/text/language"
)

// Languages is a collection of language information.
type Languages map[language.Tag]Name

// Add adds a Name to the collection. Add uses Name.Language as the key to be retrieved from the Languages.
func (v Languages) Add(name Name) {
	v[name.Language] = name
}

// Del removes a Name from Languages using the argument key.
func (v Languages) Del(key language.Tag) bool {
	delete(v, key)
	return v.Has(key)
}

// Each performs a for-each loop across Languages, executing the argument function for the Name at the current key.
func (v Languages) Each(fn func(language.Tag, Name)) {
	for k, v := range v {
		fn(k, v)
	}
}

// EachWithBreak performs a for-each loop across Languages, executing the argument function for the Name at the current key.
// Unlike Languages.Each it is possible to break out of the loop by returning true as a callback.
func (v Languages) EachWithBreak(fn func(language.Tag, Name) bool) {
	for k, v := range v {
		if fn(k, v) {
			break
		}
	}
}

// Get returns a Name from Languages using the argument key.
// Returns an additional boolean indicating whether the Name was found.
func (v Languages) Get(key language.Tag) (Name, bool) {
	name, ok := v[key]
	return name, ok
}

// Has returns a boolean indicating whether the Name exists for the argument key.
func (v Languages) Has(key language.Tag) bool {
	_, ok := v.Get(key)
	return ok
}

// Must returns a Name from Languages using the argument key.
// Unlike Languages.Get, Languages.Must panics on the condition a Name cannot be retrieved for the given key.
func (v Languages) Must(key language.Tag) Name {
	name, ok := v.Get(key)
	if !ok {
		panic(fmt.Sprintf("Languages[%s] not found", key.String()))
	}
	return name
}
