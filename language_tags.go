package nook

import (
	"slices"

	"golang.org/x/text/language"
)

var (
	supportedLanguageTags = []language.Tag{
		language.AmericanEnglish,
		language.CanadianFrench,
		language.Dutch,
		language.French,
		language.German,
		language.Italian,
		language.Japanese,
		language.Korean,
		language.LatinAmericanSpanish,
		language.Russian,
		language.SimplifiedChinese,
		language.Spanish,
		language.TraditionalChinese,
	}
)

var (
	supportedLanguageTagOrder = func() map[language.Tag]int {
		order := make(map[language.Tag]int, len(supportedLanguageTags))
		for i, tag := range supportedLanguageTags {
			order[tag] = i
		}
		return order
	}()
)

func compareLanguageTags(a, b language.Tag) int {
	aOrder, aOk := supportedLanguageTagOrder[a]
	bOrder, bOk := supportedLanguageTagOrder[b]

	switch {
	case aOk && bOk && aOrder < bOrder:
		return -1
	case aOk && bOk && aOrder > bOrder:
		return 1
	case aOk:
		return -1
	case bOk:
		return 1
	}

	switch {
	case a.String() < b.String():
		return -1
	case a.String() > b.String():
		return 1
	default:
		return 0
	}
}

func hasLanguageValue(values Languages, tag language.Tag) bool {
	name, ok := values.Get(tag)
	return ok && name.Ok()
}

// SupportedLanguageTags returns the package's canonical localized language
// coverage order.
func SupportedLanguageTags() []language.Tag {
	return append([]language.Tag(nil), supportedLanguageTags...)
}

// Complete reports whether the collection contains a non-empty localized value
// for every supported language tag.
func (v Languages) Complete() bool {
	return len(v.MissingSupportedTags()) == 0
}

// MissingSupportedTags returns the supported language tags that do not have a
// non-empty localized value, preserving canonical coverage order.
func (v Languages) MissingSupportedTags() []language.Tag {
	return v.MissingTags(supportedLanguageTags...)
}

// MissingTags returns the provided language tags that do not have a non-empty
// localized value, preserving first-seen input order.
func (v Languages) MissingTags(tags ...language.Tag) []language.Tag {
	missing := make([]language.Tag, 0, len(tags))
	seen := make(map[language.Tag]struct{}, len(tags))

	for _, tag := range tags {
		if _, ok := seen[tag]; ok {
			continue
		}
		seen[tag] = struct{}{}

		if hasLanguageValue(v, tag) {
			continue
		}

		missing = append(missing, tag)
	}

	return missing
}

// Tags returns the collection's language tags in deterministic order.
// Supported language tags are returned first in canonical coverage order and
// any additional tags follow in lexical tag-string order.
func (v Languages) Tags() []language.Tag {
	tags := make([]language.Tag, 0, len(v))
	seen := make(map[language.Tag]struct{}, len(v))

	for _, tag := range supportedLanguageTags {
		if !hasLanguageValue(v, tag) {
			continue
		}

		tags = append(tags, tag)
		seen[tag] = struct{}{}
	}

	extraTags := make([]language.Tag, 0, len(v)-len(tags))
	for tag, name := range v {
		if !name.Ok() {
			continue
		}
		if _, ok := seen[tag]; ok {
			continue
		}

		extraTags = append(extraTags, tag)
	}

	slices.SortFunc(extraTags, compareLanguageTags)
	return append(tags, extraTags...)
}
