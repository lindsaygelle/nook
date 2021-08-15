package nook

import (
	"fmt"

	"golang.org/x/text/language"
)

// Languages is a collection of language information.
type Languages map[language.Tag]Name

func (v Languages) Add(name Name) {
	v[name.Language] = name
}

func (v Languages) Del(language language.Tag) bool {
	delete(v, language)
	return v.Has(language)
}

func (v Languages) Each(fn func(language.Tag, Name)) {
	for k, v := range v {
		fn(k, v)
	}
}

func (v Languages) EachWithBreak(fn func(language.Tag, Name) bool) {
	for k, v := range v {
		if fn(k, v) {
			break
		}
	}
}

func (v Languages) Get(language language.Tag) (Name, bool) {
	name, ok := v[language]
	return name, ok
}

func (v Languages) Has(language language.Tag) bool {
	_, ok := v.Get(language)
	return ok
}

func (v Languages) Must(language language.Tag) Name {
	name, ok := v.Get(language)
	if !ok {
		panic(fmt.Sprintf("Languages[%s] not found", language.String()))
	}
	return name
}
