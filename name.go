package nook

import "golang.org/x/text/language"

type Name struct {
	Language language.Tag
	Value    string
}

func (v Name) Ok() bool {
	return v.Value != ""
}
