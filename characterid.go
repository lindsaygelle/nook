package nook

import "strings"

const characterIDSeparator = ":"

// CharacterID is a globally unique identifier for a character composed from
// its animal key and character key.
type CharacterID string

// NewCharacterID builds a global character identifier from an animal key and
// character key. Zero-value inputs produce a zero-value identifier.
func NewCharacterID(animalKey, characterKey Key) CharacterID {
	if animalKey == "" || characterKey == "" {
		return ""
	}
	return CharacterID(string(animalKey) + characterIDSeparator + string(characterKey))
}

// Ok reports whether the identifier is complete and parseable.
func (id CharacterID) Ok() bool {
	_, _, ok := id.Parts()
	return ok
}

// Parts splits a character identifier into its animal and character keys.
func (id CharacterID) Parts() (Key, Key, bool) {
	left, right, ok := strings.Cut(string(id), characterIDSeparator)
	if !ok || left == "" || right == "" {
		return "", "", false
	}
	return Key(left), Key(right), true
}

// String returns the string form of the identifier.
func (id CharacterID) String() string {
	return string(id)
}

// ID returns the globally unique identifier for the character.
func (c Character) ID() CharacterID {
	return NewCharacterID(c.Animal.Key, c.Key)
}
